package logger

import (
	"fmt"
	"os"

	"github.com/abaxoth0/Ain/logger"
	jsoniter "github.com/json-iterator/go"
)

var (
	Stdout = logger.NewStdOutLogger("")
	Stderr = logger.NewStdErrLogger("")
)

type serializer struct {
	*jsoniter.Stream
}

func newSerializer() *serializer {
	return &serializer{
		Stream: jsoniter.NewStream(jsoniter.ConfigFastest, nil, 1024),
	}
}

func (s *serializer) Reset() {
	s.Stream.Reset(nil)
}

func (s *serializer) WriteVal(v any) error {
	s.Stream.WriteVal(v)
	return nil
}

func (s *serializer) Buffer() []byte {
	return s.Stream.Buffer()
}

var DefaultConfig = &logger.FileLoggerConfig{
	Path:     "/tmp/arcturus",
	FilePerm: 0644,
	SerializerProducer: func() logger.Serializer { return newSerializer() },
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
