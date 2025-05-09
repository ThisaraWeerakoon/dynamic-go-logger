package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ThisaraWeerakoon/dynamic-go-logger/pkg/loggerfactory"

	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// Config implements the common.ConfigProvider interface
type Config struct {
	koanf *koanf.Koanf
}

func ReadFile(filename string) (*Config, error) {
	k := koanf.New(".")
	f := file.Provider(filename)
	if err := k.Load(f, toml.Parser()); err != nil {
		return nil, err
	}
	cfg := &Config{
		koanf: k,
	}
	return cfg, nil
}

func (c *Config) IsSet(key string) bool {
	return c.koanf.Exists(key)
}

func (c *Config) Watch(ctx context.Context, filename string) {
	f := file.Provider(filename)

	f.Watch(func(event interface{}, err error) {
		if err != nil {
			log.Printf("watch error: %v", err)
			return
		}
		// Throw away the old config and load a fresh copy.
		log.Println("config changed. Reloading ...")
		new_k := koanf.New(".")
		if err := new_k.Load(f, toml.Parser()); err != nil {
			log.Printf("error loading new config: %v", err)
			return
		}
		// Update the config
		c.koanf = new_k

		// Update the logger configuration
		var levelMap map[string]string
		var slogHandlerConfig loggerfactory.SlogHandlerConfig

		c.MustUnmarshal("logger.level.packages", &levelMap)
		c.MustUnmarshal("logger.handler", &slogHandlerConfig)

		cm := loggerfactory.GetConfigManager()
		cm.SetLogLevelMap(&levelMap)
		cm.SetSlogHandlerConfig(slogHandlerConfig)
	})
}

func (c *Config) Unmarshal(key string, out interface{}) error {
	err := c.koanf.Unmarshal(key, out)
	if err != nil {
		return fmt.Errorf("cannot unmarshal config for key %q: %v", key, err)
	}
	return nil
}

func (c *Config) MustUnmarshal(key string, out interface{}) {
	err := c.Unmarshal(key, out)
	if err != nil {
		panic(err)
	}
}

func InitializeConfig(confFolderPath string) error {
	files, err := os.ReadDir(confFolderPath)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return fmt.Errorf("no config files found in %s", confFolderPath)
	}
	// {"LoggerConfig", "ServerConfig"}
	for _, configurationType := range []string{"LoggerConfig"} {
		configFilePath := filepath.Join(confFolderPath, configurationType+".toml")
		cfg, err := ReadFile(configFilePath)
		if err != nil {
			return fmt.Errorf("cannot read config file: %w", err)
		}

		switch configurationType {
		case "LoggerConfig":
			var levelMap map[string]string
			var slogHandlerConfig loggerfactory.SlogHandlerConfig

			if cfg.IsSet("logger") {
				cfg.MustUnmarshal("logger.handler", &slogHandlerConfig)
				cfg.MustUnmarshal("logger.level.packages", &levelMap)
			}

			cm := loggerfactory.GetConfigManager()
			cm.SetLogLevelMap(&levelMap)
			cm.SetSlogHandlerConfig(slogHandlerConfig)

			// Start watching for config changes
			cfg.Watch(context.Background(), configFilePath)

			// Add the config to the context
			// configContext := ctx.Value(utils.ConfigContextKey).(*artifacts.ConfigContext)
			// configContext.AddLoggerConfig(cfg)

			// case "ServerConfig":
			// TODO: Add server config
		}
	}
	return nil
}
