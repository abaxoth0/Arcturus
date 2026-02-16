package config

import "flag"

var (
	DebugMode bool
	TraceLogs bool
	Verbose  bool
	LogPath	  string
)

func Parse() {
	debug := flag.Bool("d", false, "Debug - Run app in debug mode")
	traceLogs := flag.Bool("t", false, "Trace - Enable trace logs")
	verbose := flag.Bool("v", false, "Verbose - Show logs")
	logPath := flag.String("logs-dir", "/tmp/arcturus", "Path to the logs directory")

	flag.Parse()

	DebugMode = *debug
	TraceLogs = *traceLogs
	Verbose = *verbose
	LogPath = *logPath
}
