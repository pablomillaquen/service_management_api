package logger

import (
	"log"
	"os"
	"time"
)

type Level int

const (
	LevelInfo  Level = iota
	LevelWarn  Level = iota
	LevelError Level = iota
)

func (l Level) String() string {
	switch l {
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

var logger *log.Logger

func Init() {
	logger = log.New(os.Stdout, "", 0)
	Info("Logger initialized")
}

func logf(level Level, format string, args ...interface{}) {
	if logger == nil {
		logger = log.New(os.Stdout, "", 0)
	}
	prefix := time.Now().Format(time.RFC3339) + " [" + level.String() + "] "
	logger.Printf(prefix+format, args...)
}

func Info(format string, args ...interface{}) {
	logf(LevelInfo, format, args...)
}

func Warn(format string, args ...interface{}) {
	logf(LevelWarn, format, args...)
}

func Error(format string, args ...interface{}) {
	logf(LevelError, format, args...)
}
