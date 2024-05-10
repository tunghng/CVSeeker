package main

import (
	"CVSeeker/cmd/CVSeeker/internal/providers"
	_ "CVSeeker/docs"
	"CVSeeker/internal/ginServer"
	"CVSeeker/pkg/cfg"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

// @title           CVSeeker Server
// @version         1.0
// @description     This is the server for api endpoints related to the CVSeeker application
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	providers.BuildContainer()

	if os.Getenv("ENVIRONMENT") == cfg.EnvironmentLocal {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	log.Println("Preparing and running main application . . . !")
	if err := run(); err != nil {
		log.Fatalf("Running HTTP server: %v", err)
	}
}

func run() error {
	c := providers.GetContainer()
	if c == nil {
		log.Fatalf("Container hasn't been initialized yet")
	}
	var s ginServer.Server
	if err := c.Invoke(func(_s ginServer.Server) { s = _s }); err != nil {
		return err
	}

	if err := s.Open(); err != nil {
		return err
	}

	return nil
}
