package slog

import (
	"log"
	"os"
)

var (
	// VerboseLog is whethere allow verbose log output.
	VerboseLog = false
)

// SimpleLogger is a simple logger.
type SimpleLogger interface {
	Debugf(formatString string, args ...interface{})
	Infof(formatString string, args ...interface{})
	Fatalf(formatString string, args ...interface{})
}

type simpleLogger struct {
	logger *log.Logger
}

// New creates a new logger.
func New(prefix string) SimpleLogger {
	l := &simpleLogger{}
	l.logger = log.New(os.Stderr, prefix, log.Lshortfile)
	return l
}

func (s *simpleLogger) Debugf(formatString string, args ...interface{}) {
	if VerboseLog {
		s.logger.Printf(formatString, args...)
	}
}

func (s *simpleLogger) Infof(formatString string, args ...interface{}) {
	s.logger.Printf(formatString, args...)
}

func (s *simpleLogger) Fatalf(formatString string, args ...interface{}) {
	s.logger.Fatalf(formatString, args...)
}
