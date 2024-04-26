package ginMiddleware

import (
	"github.com/gin-gonic/gin"
	"grabber-match/internal/ginLogger"
	"log"
	"net/http"
	"runtime/debug"
)

// Recovery handles the request when it panic.
func Recovery() gin.HandlerFunc {
	panicPayload := gin.H{
		"meta": gin.H{
			"code":    http.StatusServiceUnavailable,
			"message": "Dịch vụ xảy ra lỗi khi xử lý yêu cầu",
		},
	}
	return func(c *gin.Context) {
		defer func() {
			defer func() {
				if rec := recover(); rec != nil {
					log.Println(string(debug.Stack()))
					ginLogger.Gin(c).Errorf("Stack trace panic: %v", string(debug.Stack()))
					c.JSON(http.StatusInternalServerError, panicPayload)
					c.Abort()
				}
			}()
			c.Next()
		}()
	}
}
