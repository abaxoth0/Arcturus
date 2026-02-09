package logger

import (
	"fmt"
	"os"

	"github.com/abaxoth0/Ain/logger"
)

var (
	Stdout = logger.NewStdOutLogger("")
	Stderr = logger.NewStdErrLogger("")
)

var DefaultConfig = &logger.FileLoggerConfig{
	Path:     "/tmp/arcturus",
	FilePerm: 0744,
	LoggerConfig: &logger.LoggerConfig{
		ApplicationName: "arcturus",
	},
}

var Default = func() *logger.FileLogger {
	log, err := logger.NewFileLogger(DefaultConfig)
	if err != nil {
		fmt.Println("Failed to setup default Logger:", err)
		os.Exit(1)
	}
	return log
}()
