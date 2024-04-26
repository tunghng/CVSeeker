package providers

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	_swaggerFiles "github.com/swaggo/files"
	_ginSwagger "github.com/swaggo/gin-swagger"
	"grabber-match/cmd/grabber-match/internal/cfg"
	"grabber-match/cmd/grabber-match/internal/handlers"
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

		//faQuiz := baseRoute.Group("/faquiz")
		//v1FaQuiz := faQuiz.Group("/v1")

	}
}
