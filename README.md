
# go-logger

The **go-logger** repository demonstrates a dynamic (hot deployable) package-level logger implementation in Go using the standard `log/slog` library. The project shows how to update log levels at runtime by leveraging a configuration file and the observer pattern.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [How It Works](#how-it-works)
- [Contributing](#contributing)
- [License](#license)

## Overview

In production systems, you might need to adjust logging levels on the fly for better debugging and performance tuning. This repository offers an example of:
- Using a configuration file (`LoggerConfig.toml`) to set log levels per package.
- Dynamically updating logger settings at runtime without restarting the application.
- Implementing the observer pattern to propagate configuration changes to all logger users.

## Features

- **Dynamic Logger Configuration:** Adjust logger levels and handler settings at runtime.
- **Observer Pattern:** Components automatically update their loggers when configuration changes.
- **Package-Level Logging:** Different logging levels can be applied to individual packages.
- **Configuration Management:** Uses [Koanf](https://github.com/knadh/koanf) for configuration parsing and live-reloading.

## Project Structure

```plaintext
go-logger/
├── cmd/
│   └── myapp/
│       └── myapp.go       # Main application entry point.
├── conf/
│   └── LoggerConfig.toml  # Configuration file for logger settings.
├── internal/
│   └── pkg/
│       ├── packageA/
│       │   └── programA.go   # Package A logging implementation.
│       ├── packageB/
│       │   └── programB.go   # Package B logging implementation.
│       └── packageC/
│           └── packageC.go   # Package C logging implementation.
├── pkg/
│   ├── config/
│   │   └── config.go         # Configuration loading and watching.
│   └── loggerfactory/
│       └── loggerfactory.go  # Logger factory and configuration management.
├── go.mod                  # Go module file.
└── go.sum                  # Go dependency checksum file.
```

## Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/YourUsername/go-logger.git
   cd go-logger
   ```

2. **Build the Application:**

   ```bash
   go build ./cmd/myapp
   ```

3. **Run the Application:**

   ```bash
   ./myapp
   ```

The application will continuously print logs every 5 seconds. You can modify the configuration file to see dynamic updates to log levels.

## Configuration

The logger is configured using a TOML file located in the `conf` folder. Below is a sample configuration (`LoggerConfig.toml`):

```toml
[logger]
level.default = "warn"

[logger.level.packages]
packageA = "warn"
packageB = "info"
packageC = "debug"

[logger.handler]
format = "json"
outputPath = "stdout"
```

- **logger.level.packages:** Defines the log level for each package.
- **logger.handler:** Specifies the format and output for the logs (e.g., JSON to stdout).

Any changes saved to this file trigger a reload, and the new settings are applied at runtime.

## Usage

In your main application (see `cmd/myapp/myapp.go`), the configuration is initialized, and logger instances for various packages are created using the logger factory:

```go
func main() {
    // Determine configuration file path.
    _, filename, _, ok := runtime.Caller(0)
    if !ok {
        log.Fatal("Could not get current file path")
    }
    currentDir := filepath.Dir(filename)
    confPath := filepath.Join(filepath.Dir(currentDir), "..", "conf")

    // Initialize configuration and start watching for changes.
    if err := config.InitializeConfig(confPath); err != nil {
        log.Fatalf("Initialization error: %s", err.Error())
    }

    // Create components and output logs.
    a := packageA.NewA("value1", 42)
    b := packageB.NewB("value2", 84)
    for {
        fmt.Println("=============================================")
        a.ShowLogs()
        b.ShowLogs()
        b.InitiatePackageC() // Dynamically triggers packageC logging.
        time.Sleep(5 * time.Second)
    }
}
```

Each package implements an `UpdateLogger()` method that is automatically called when the configuration changes, ensuring that every component always uses the correct logging settings.

## How It Works

1. **Configuration Loading and Watching:**  
   The `pkg/config/config.go` file uses Koanf to load the initial configuration from `LoggerConfig.toml` and watches the file for changes. When a change is detected, the new configuration is loaded and pushed to the logger factory via the configuration manager.

2. **Observer Pattern:**  
   The `ConfigManager` in `pkg/loggerfactory/loggerfactory.go` maintains a registry of components (observers) that use loggers. When configuration changes occur, `ConfigManager` notifies each registered component via its `UpdateLogger()` method.

3. **Logger Factory:**  
   The `GetLogger` function retrieves logger configuration based on the package name and returns a new logger instance with the appropriate log level and handler.

This modular design ensures that logging remains flexible and can be adjusted on the fly, which is especially useful for debugging production issues without downtime.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvement, please open an issue or submit a pull request.

1. Fork the repository.
2. Create your feature branch: `git checkout -b feature/my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin feature/my-new-feature`
5. Create a new Pull Request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
