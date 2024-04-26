package ginMiddleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	newrelic "github.com/newrelic/go-agent"
	"grabber-match/internal/dtos"
)

type key int

const (
	KeyNrID key = iota
)

// RequestIDLoggingMiddleware RequestIDLoggingMiddleware
func RequestIDLoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqID string
		reqID = c.GetHeader(dtos.HeaderXRequestID)
		if reqID == "" {
			reqID = uuid.New().String()
			c.Header(dtos.HeaderXRequestID, reqID)
		}
		c.Next()
	}
}

// SetNewRelicInContext get the request context populated
func SetNewRelicInContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Setup context
		ctx := c.Request.Context()

		//Set newrelic context
		var txn newrelic.Transaction
		//newRelicTransaction is the key populated by nrgin Middleware
		value, exists := c.Get("newRelicTransaction")
		if exists {
			if v, ok := value.(newrelic.Transaction); ok {
				txn = v
			}
			ctx = context.WithValue(ctx, KeyNrID, txn)
		}
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
