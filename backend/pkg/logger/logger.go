package logger

import (
	"context"
	"log"
	"os"
	"strings"
)

// TraceIDField TraceIDField
const TraceIDField = "trace_id"

// MessageIDField MessageIDField
const MessageIDField = "message_id"

// ContextTraceID ContextTraceID
const ContextTraceID = "ctxTraceId"

// Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

const (
	//Debug has verbose message
	Debug = "debug"
	//Info is default log level
	Info = "info"
	//Warn is for logging messages about possible issues
	Warn = "warn"
	//Error is for logging errors
	Error = "error"
	//Fatal is for logging fatal messages. The sytem shutsdown after logging the message.
	Fatal = "fatal"
)

// Logger is our contract for the logger
type Logger interface {
	Debugf(format string, args ...interface{})

	Infof(format string, args ...interface{})

	Warnf(format string, args ...interface{})

	Errorf(format string, args ...interface{})

	Fatalf(format string, args ...interface{})

	Panicf(format string, args ...interface{})

	TraceCtx(ctx context.Context) Logger

	WithFields(keyValues Fields) Logger
}

// Configuration stores the config for the logger
// For some loggers there can only be one level across writers, for such the level of Console is picked by default
type Configuration struct {
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      string
	EnableFile        bool
	FileJSONFormat    bool
	FileLevel         string
	FileLocation      string
}

func NewLogger() Logger {
	loggerLevel := strings.ToLower(os.Getenv("LOG_LEVEL"))
	if loggerLevel == "" {
		loggerLevel = Debug
	}

	config := Configuration{
		EnableConsole:     true,
		ConsoleLevel:      loggerLevel,
		ConsoleJSONFormat: true,
	}
	logger, err := newGCPZapLogger(config)
	if err != nil {
		log.Fatal(err)
	}
	return logger
}
