package providers

import (
	"grabber-match/cmd/grabber-match/internal/cfg"
	"grabber-match/internal/dtos"
	"grabber-match/internal/errors"
	"grabber-match/internal/ginServer"
	"grabber-match/internal/meta"
	"grabber-match/pkg/api"
	"grabber-match/pkg/app"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// newServerConfig returns a *server.Config.
func newServerConfig() *ginServer.Config {
	return &ginServer.Config{
		Addr: viper.GetString(cfg.ConfigKeyHttpAddress),
		Port: viper.GetInt64(cfg.ConfigKeyHttpPort),
	}
}

func newErrorParserConfig() *errors.ErrorParserConfig {
	staticErrorCfgPath := "./statics/errors.toml"
	return &errors.ErrorParserConfig{PathConfigError: staticErrorCfgPath}
}

func newGinEngine() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, dtos.Response{
			Meta: meta.Meta{
				Code:    http.StatusNotFound,
				Message: "Page not found",
			}})
	})

	return r
}

//------------------------------------------------------------------

func newApiConfig() *api.Config {
	return &api.Config{
		DefaultPageSize: viper.GetInt64(cfg.ConfigApiDefaultPageSize),
		MinPageSize:     viper.GetInt64(cfg.ConfigApiMinPageSize),
		MaxPageSize:     viper.GetInt64(cfg.ConfigApiMaxPageSize),
	}
}

func newAppConfig() *app.Config {
	return &app.Config{
		DirectoryTemp:    viper.GetString(cfg.ConfigKeyFolderTmp),
		BucketName:       viper.GetString(cfg.ConfigKeyGCSBucket),
		CdnBucketName:    viper.GetString(cfg.ConfigKeyGCSBucketCDN),
		CdnRootFolder:    viper.GetString(cfg.ConfigKeyGCSBucketCDNRootFolder),
		CdnURL:           viper.GetString(cfg.ConfigKeyCDNUrl),
		URLGoogleStorage: viper.GetString(cfg.URLGoogleStorage),
	}
}

// LoadConfigEnv loads configuration from the given list of paths and populates it into the Config variable.
func newCfgReader() *viper.Viper {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return v
}
