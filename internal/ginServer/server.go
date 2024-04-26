package ginServer

import (
	"errors"
	"fmt"
)

// Server describes http server.
type Server interface {
	Open() error
	Close()
}

// Errors definition.
var (
	ErrNilRoutingFn = errors.New("routing function cannot be nil")
)

// Config represents server configuration.
type Config struct {
	Addr string // 0.0.0.0
	Port int64  // 8080
}

// ListenerAddr returns HTTP server address.
func (_this *Config) ListenerAddr() string {
	var (
		addr = "0.0.0.0"
		port = int64(1323)
	)
	if _this.Addr != "" {
		addr = _this.Addr
	}
	if _this.Port > 0 && _this.Port < 65535 {
		port = _this.Port
	}
	return fmt.Sprintf("%v:%v", addr, port)
}
