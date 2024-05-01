// Package logger ...
package ginLogger

import (
	"CVSeeker/internal/dtos"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimestampFieldName = "_datetime"
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	zerolog.MessageFieldName = "handler_func"
}

// MiddlewareGin is middleware function for Gin.
func MiddlewareGin(appName string, logLevel zerolog.Level) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader(HeaderXRequestID)
		if reqID == "" {
			reqID = uuid.New().String()
		}
		var (
			startTime      = time.Now()
			reqEndpoint    = c.Request.URL.Path
			reqQueries     = c.Request.URL.RawQuery
			reqMethod      = c.Request.Method
			handlerName    = c.HandlerName()
			xForwardedFor  = c.GetHeader(HeaderXForwardedFor)
			xRequestId     = reqID
			trueClientIP   = c.GetHeader(HeaderTrueClientIP)
			cfConnectingIP = c.GetHeader(HeaderCFConnectingIP)
			deviceID       = c.GetHeader(HeaderDeviceID)
		)

		if strings.Contains(reqEndpoint, "health-check") || strings.Contains(reqEndpoint, "swagger") {
			return
		}

		zrLogger := log.With().
			Str("app_name", appName).
			Str("endpoint", reqEndpoint).
			Str("method", reqMethod).
			Str(DeviceID, deviceID).
			Str(XForwardedFor, xForwardedFor).
			Str(XRequestID, xRequestId).
			Str(TrueClientIP, trueClientIP).
			Str(CFConnectingIP, cfConnectingIP).
			Logger()
		if reqQueries != "" {
			zrLogger = zrLogger.With().Str("queries", reqQueries).Logger()
		}

		c.Next()

		logger := &logger{
			GinCtx:  c,
			EchoCtx: nil,
		}

		for _, key := range []string{ErrKey, InfoKey, DebugKey, WarningKey, CustomDataKey} {
			zrLogger = prepareLogData(logger, zrLogger, key, logLevel)
		}

		var (
			finishTime        = time.Now()
			latency           = finishTime.Sub(startTime)
			statusCode        = c.Writer.Status()
			response, existed = c.Get(dtos.ContextResponse)
		)
		if existed {
			zrLogger.Log().
				Str("severity", getSeverityLevelLog(logger, logLevel)).
				Str("message", fmt.Sprintf("[%s] [%s] %13v %d %s",
					xRequestId, reqMethod, latency, statusCode, reqEndpoint)).
				Dur("latency", latency).
				Int("status_code", statusCode).
				RawJSON("response", response.([]byte)).
				Msg(handlerName)
		} else {
			zrLogger.Log().
				Str("severity", getSeverityLevelLog(logger, logLevel)).
				Str("message", fmt.Sprintf("[%s] [%s] %13v %d %s",
					xRequestId, reqMethod, latency, statusCode, reqEndpoint)).
				Dur("latency", latency).
				Int("status_code", statusCode).
				Msg(handlerName)
		}
	}
}

func getSeverityLevelLog(logger *logger, logLevelDefault zerolog.Level) string {
	var currentLogLevel = logLevelDefault
	if logData := logger.GetLogData(DebugKey); len(logData) != 0 && zerolog.DebugLevel >= logLevelDefault {
		currentLogLevel = zerolog.DebugLevel
	}

	if logData := logger.GetLogData(InfoKey); len(logData) != 0 && zerolog.InfoLevel >= logLevelDefault {
		currentLogLevel = zerolog.InfoLevel
	}

	if logData := logger.GetLogData(WarningKey); len(logData) != 0 && zerolog.WarnLevel >= logLevelDefault {
		currentLogLevel = zerolog.WarnLevel
	}

	if logData := logger.GetLogData(ErrKey); len(logData) != 0 && zerolog.ErrorLevel >= logLevelDefault {
		currentLogLevel = zerolog.ErrorLevel
	}

	switch currentLogLevel {
	case zerolog.DebugLevel:
		return "DEBUG"
	case zerolog.InfoLevel:
		return "INFO"
	case zerolog.WarnLevel:
		return "WARNING"
	case zerolog.ErrorLevel:
		return "ERROR"
	}
	return "DEFAULT"
}

func prepareLogData(logger *logger, zrLogger zerolog.Logger, key string, currentLogLevel zerolog.Level) zerolog.Logger {
	logData := logger.GetLogData(key)
	switch key {
	case CustomDataKey:
		for _, key := range logData {
			zrLogger = zrLogger.With().Interface(key, logger.Get(key)).Logger()
		}
	case ErrKey:
		if zerolog.ErrorLevel >= currentLogLevel {
			for i, data := range logData {
				zrLogger = zrLogger.With().Str(fmt.Sprintf("%v_%v", key, i), data).Logger()
			}
		}
	case InfoKey:
		if zerolog.InfoLevel >= currentLogLevel {
			for i, data := range logData {
				zrLogger = zrLogger.With().Str(fmt.Sprintf("%v_%v", key, i), data).Logger()
			}
		}
	case DebugKey:
		if zerolog.DebugLevel >= currentLogLevel {
			for i, data := range logData {
				zrLogger = zrLogger.With().Str(fmt.Sprintf("%v_%v", key, i), data).Logger()
			}
		}
	case WarningKey:
		if zerolog.WarnLevel >= currentLogLevel {
			for i, data := range logData {
				zrLogger = zrLogger.With().Str(fmt.Sprintf("%v_%v", key, i), data).Logger()
			}
		}
	default:
		for i, data := range logData {
			zrLogger = zrLogger.With().Str(fmt.Sprintf("%v_%v", key, i), data).Logger()
		}
	}
	return zrLogger
}
