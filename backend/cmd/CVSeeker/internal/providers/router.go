package providers

import (
	"CVSeeker/cmd/CVSeeker/internal/cfg"
	"CVSeeker/cmd/CVSeeker/internal/handlers"
	"CVSeeker/internal/ginLogger"
	commonMiddleware "CVSeeker/internal/ginMiddleware"
	"CVSeeker/internal/ginServer"
	"CVSeeker/pkg/websocket"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	_swaggerFiles "github.com/swaggo/files"
	_ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

// setupRouter setup router.
func setupRouter(hs *handlers.Handlers) ginServer.GinRoutingFn {
	return func(router *gin.Engine) {
		// CORS configuration
		corsConfig := cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}

		router.Use(
			cors.New(corsConfig),
			gzip.Gzip(gzip.DefaultCompression),
			commonMiddleware.RequestIDLoggingMiddleware(),
			ginLogger.MiddlewareGin(AppName, zerolog.InfoLevel),
			commonMiddleware.Recovery(),
		)

		baseRoute := router.Group(viper.GetString(cfg.ConfigKeyContextPath))
		baseRoute.GET("swagger/*any", _ginSwagger.WrapHandler(_swaggerFiles.Handler))

		data := baseRoute.Group("/resumes")
		{
			data.POST("/upload", hs.DataProcessingHandler.ProcessDataHandler())
			data.GET("/upload", hs.DataProcessingHandler.GetAllUploadsHandler())
			data.POST("/batch/upload", hs.DataProcessingHandler.ProcessDataBatchHandler())

			data.POST("/search", hs.SearchHandler.HybridSearch())
			data.GET("/:id", hs.SearchHandler.GetDocumentByID())
			data.DELETE("/:id", hs.SearchHandler.DeleteDocumentByID())

			data.POST("/thread/start", hs.ChatbotHandler.StartChatSession())
			data.POST("/thread/:threadId/send", hs.ChatbotHandler.SendMessage())
			data.GET("/thread/:threadId/messages", hs.ChatbotHandler.ListMessage())
			data.GET("/thread", hs.ChatbotHandler.GetAllThreads())
			data.GET("/thread/:threadId", hs.ChatbotHandler.GetResumesByThreadID())
			data.POST("/thread/:threadId/updateName", hs.ChatbotHandler.UpdateThreadName())
		}

		router.GET("/ws", func(c *gin.Context) {
			// Error handling omitted for brevity
			_, err := websocket.HandleWebSocket(c.Writer, c.Request)
			if err != nil {
				// Log error or handle it
				return
			}
		})
	}
}
