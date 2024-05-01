package ginLogger

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo"
)

/* How to use logger Gin

logger.Gin(c).Errorf("Errorf message")
logger.Gin(c).Infof("Infof")

*/

// Logger key constants definition.
const (
	ErrKey         = "_err"
	InfoKey        = "_info"
	DebugKey       = "_debug"
	WarningKey     = "_warning"
	CustomDataKey  = "_custom_data"
	CFConnectingIP = "cf_connecting_ip"
	XForwardedFor  = "x_forwarded_for"
	TrueClientIP   = "true_client_ip"
	XKeyCode       = "x_key_code"
	XRequestID     = "x_request_id"
	DeviceID       = "device_id"
)

const (
	HeaderXRequestID     = "X-Request-ID"
	HeaderCFConnectingIP = "CF-Connecting-IP"
	HeaderXForwardedFor  = "X-Forwarded-For"
	HeaderTrueClientIP   = "True-Client-IP"
	HeaderDeviceID       = "sab-device-id"
)

// LogArr is array of string.
type LogArr []string

// Gin logs with Gin framework Context.
func Gin(c *gin.Context) Logger {
	return &logger{
		GinCtx:  c,
		EchoCtx: nil,
	}
}

// Echo logs with Echo framework Context.
func Echo(c echo.Context) Logger {
	return &logger{
		GinCtx:  nil,
		EchoCtx: c,
	}
}

// Logger is the best logging lib.
type Logger interface {
	With(key string, value interface{}) Logger
	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

type logger struct {
	GinCtx  *gin.Context
	EchoCtx echo.Context
}

func (_this *logger) With(key string, value interface{}) Logger {
	if value != nil {
		_this.Append(CustomDataKey, key)
		_this.Set(key, value)
	}
	return _this
}

func (_this *logger) Errorf(format string, args ...interface{}) {
	_this.Append(ErrKey, format, args...)
}

func (_this *logger) Infof(format string, args ...interface{}) {
	_this.Append(InfoKey, format, args...)
}

func (_this *logger) Warningf(format string, args ...interface{}) {
	_this.Append(WarningKey, format, args...)
}

func (_this *logger) Debugf(format string, args ...interface{}) {
	_this.Append(DebugKey, format, args...)
}

func (_this *logger) Set(key string, value interface{}) {
	switch {
	case _this.GinCtx != nil:
		_this.GinCtx.Set(key, value)
	case _this.EchoCtx != nil:
		_this.EchoCtx.Set(key, value)
	}
}

func (_this *logger) Get(key string) interface{} {
	switch {
	case _this.GinCtx != nil:
		value, _ := _this.GinCtx.Get(key)
		return value
	case _this.EchoCtx != nil:
		return _this.EchoCtx.Get(key)
	}
	return nil
}

func (_this *logger) GetLogData(key string) LogArr {
	val := _this.Get(key)
	if val != nil {
		return val.(LogArr)
	}
	return nil
}

func (_this *logger) Initial(key string) {
	switch {
	case _this.GinCtx != nil:
		val, exists := _this.GinCtx.Get(key)
		if !exists || val == nil {
			_this.GinCtx.Set(key, LogArr{})
		}
	case _this.EchoCtx != nil:
	}
}

func (_this *logger) Append(key string, format string, args ...interface{}) {
	_this.Initial(key)
	if value := _this.Get(key); value != nil {
		logArr := value.(LogArr)
		logArr = append(logArr, fmt.Sprintf(format, args...))
		_this.Set(key, logArr)
	}
}
