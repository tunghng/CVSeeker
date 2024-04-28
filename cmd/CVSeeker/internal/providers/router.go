package providers

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	_swaggerFiles "github.com/swaggo/files"
	_ginSwagger "github.com/swaggo/gin-swagger"
	"grabber-match/cmd/CVSeeker/internal/cfg"
	"grabber-match/cmd/CVSeeker/internal/handlers"
	"grabber-match/internal/ginLogger"
	commonMiddleware "grabber-match/internal/ginMiddleware"
	"grabber-match/internal/ginServer"
)

// setupRouter setup router.
func setupRouter(hs *handlers.Handlers) ginServer.GinRoutingFn {
	return func(router *gin.Engine) {

		router.Use(
			gzip.Gzip(gzip.DefaultCompression),
			commonMiddleware.RequestIDLoggingMiddleware(),
			ginLogger.MiddlewareGin(AppName, zerolog.InfoLevel),
			commonMiddleware.Recovery(),
		)

		baseRoute := router.Group(viper.GetString(cfg.ConfigKeyContextPath))
		baseRoute.GET("swagger/*any", _ginSwagger.WrapHandler(_swaggerFiles.Handler))

		data := baseRoute.Group("/resumes")
		{
			data.POST("", hs.DataProcessingHandler.ProcessDataHandler())
		}
	}
}
