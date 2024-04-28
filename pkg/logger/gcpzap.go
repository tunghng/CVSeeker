package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"reflect"
)

type gcpZapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

func newGCPZapLogger(config Configuration) (Logger, error) {
	var LogLevelCfg zapcore.Level
	switch config.ConsoleLevel {
	case Debug:
		LogLevelCfg = zapcore.DebugLevel
	case Info:
		LogLevelCfg = zapcore.InfoLevel
	case Warn:
		LogLevelCfg = zapcore.WarnLevel
	case Error:
		LogLevelCfg = zapcore.ErrorLevel
	case Fatal:
		LogLevelCfg = zapcore.FatalLevel
	default:
		LogLevelCfg = zapcore.InfoLevel
	}

	fluentdConfig, _ := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(LogLevelCfg),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "message",
			StacktraceKey: "stacktrace",
			LevelKey:      "severity",
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			TimeKey:       "time",
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			CallerKey:     "caller",
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
	}.Build()

	return &gcpZapLogger{
		sugaredLogger: fluentdConfig.Sugar(),
	}, nil
}

func (l *gcpZapLogger) Debugf(format string, args ...interface{}) {
	l.sugaredLogger.Debugf(format, args...)
}

func (l *gcpZapLogger) Infof(format string, args ...interface{}) {
	l.sugaredLogger.Infof(format, args...)
}

func (l *gcpZapLogger) Warnf(format string, args ...interface{}) {
	l.sugaredLogger.Warnf(format, args...)
}

func (l *gcpZapLogger) Errorf(format string, args ...interface{}) {
	l.sugaredLogger.Errorf(format, args...)
}

func (l *gcpZapLogger) Fatalf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

func (l *gcpZapLogger) Panicf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

func (l *gcpZapLogger) WithFields(fields Fields) Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.sugaredLogger.With(f...)
	return &zapLogger{newLogger}
}

func (l *gcpZapLogger) TraceCtx(ctx context.Context) Logger {
	if ctx != nil && reflect.TypeOf(ctx).String() != "*context.emptyCtx" {
		return l.WithFields(Fields{TraceIDField: ctx.Value(ContextTraceID)})
	}
	return l
}
