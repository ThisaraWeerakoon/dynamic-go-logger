package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"time"

	"github.com/ThisaraWeerakoon/dynamic-go-logger/pkg/config"

	"github.com/ThisaraWeerakoon/dynamic-go-logger/internal/pkg/packageA"
	"github.com/ThisaraWeerakoon/dynamic-go-logger/internal/pkg/packageB"
)

func main() {
	// Get the absolute path of the conf folder
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Could not get current file path")
	}

	// Navigate from current file to conf directory
	// Go from /cmd/myapp/myapp.go to /cmd/conf/
	currentDir := filepath.Dir(filename)
	confPath := filepath.Join(filepath.Dir(currentDir),"..", "conf")

	errConfig := config.InitializeConfig(confPath)
	if errConfig != nil {
		log.Fatalf("Initialization error: %s", errConfig.Error())
	}

	a := packageA.NewA("value1", 42)
	b := packageB.NewB("value2", 84)
	for{
		fmt.Println("=============================================")
		a.ShowLogs()
		b.ShowLogs()
		b.InitiatePackageC()
		time.Sleep(5 * time.Second)
	}

}
