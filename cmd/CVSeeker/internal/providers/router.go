package providers

import (
	"CVSeeker/cmd/CVSeeker/internal/cfg"
	"CVSeeker/cmd/CVSeeker/internal/handlers"
	"CVSeeker/internal/ginLogger"
	commonMiddleware "CVSeeker/internal/ginMiddleware"
	"CVSeeker/internal/ginServer"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	_swaggerFiles "github.com/swaggo/files"
	_ginSwagger "github.com/swaggo/gin-swagger"
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

		search := baseRoute.Group("/search")
		{
			search.GET("", hs.SearchHandler.HybridSearchHandler())
		}
	}
}
