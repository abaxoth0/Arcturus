package config

import "flag"

var (
	DebugMode bool
	TraceLogs bool
	ShowLogs  bool
	LogPath	  string
)

func Parse() {
	debug := flag.Bool("debug", false, "Run app in debug mode")
	traceLogs := flag.Bool("trace", false, "Enable trace logs")
	showLogs := flag.Bool("log", false, "Show logs in terminal")
	logPath := flag.String("log-path", "/tmp/arcturus", "Path to the logs directory")

	flag.Parse()

	DebugMode = *debug
	TraceLogs = *traceLogs
	ShowLogs = *showLogs
	LogPath = *logPath
}
