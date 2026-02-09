package main

import (
	"arcturus/packages/lexer"
	log "arcturus/packages/shared/logger"
	"arcturus/packages/shared/config"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/abaxoth0/Ain/logger"
)

var appLogger = logger.NewSource("APP", log.Stdout)

func main() {
	config.Parse()

	log.DefaultConfig.Debug = config.DebugMode
	log.DefaultConfig.Trace = config.TraceLogs
	log.DefaultConfig.Path = config.LogPath

	if err := log.Default.Init(); err != nil {
		appLogger.Fatal("Failed to initialize logger", err.Error(), nil)
	}
	if config.ShowLogs {
		if err := log.Default.AddForwarding(log.Stdout); err != nil {
			appLogger.Error("Failed to add Default -> Stdout logs forwarding", err.Error(), nil)
		}
	}

	go func() {
		if err := log.Default.Start(); err != nil {
			appLogger.Fatal("Failed to start logger", err.Error(), nil)
		}
	}()
	defer func() {
		if err := log.Default.Stop(true); err != nil {
			appLogger.Error("Failed to stop logger", err.Error(), nil)
		}
	}()

	// Reserve some time for logger to start up
	time.Sleep(time.Millisecond * 50)

	file, err := os.OpenFile("example.arc", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	tokens, err := lexer.Parse(content)
	if err != nil {
		appLogger.Fatal("Failed to parse input", err.Error(), nil)
	}

	for _, tk := range tokens {
		r := tk.String()
		fmt.Printf("%s ", r)
	}

	println()
}
