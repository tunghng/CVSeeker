package providers

import (
	"go.uber.org/dig"
	"grabber-match/cmd/CVSeeker/internal/handlers"
	services "grabber-match/cmd/CVSeeker/internal/service"
	"grabber-match/internal/errors"
	"grabber-match/internal/ginServer"
	commonHandler "grabber-match/internal/handlers"
	"grabber-match/internal/repositories"
	"grabber-match/pkg/cfg"
	"grabber-match/pkg/elasticsearch"
	"grabber-match/pkg/gpt"
	"grabber-match/pkg/logger"
)

const (
	// AppName - name of module
	AppName = "CVSeeker"
)

func init() {
	cfg.SetupConfig()
}

// container is a global Container.
var container *dig.Container

// BuildContainer build all necessary containers.
func BuildContainer() *dig.Container {
	container = dig.New()
	{
		_ = container.Provide(newCfgReader)
		_ = container.Provide(newApiConfig)
		_ = container.Provide(newAppConfig)
		_ = container.Provide(newGinEngine)
		_ = container.Provide(setupRouter)
		_ = container.Provide(newServerConfig)
		_ = container.Provide(newErrorParserConfig)
		_ = container.Provide(newMySQLConnection, dig.Name("talentAcquisitionDB"))

		_ = container.Provide(logger.NewLogger)
		_ = container.Provide(errors.NewErrorParser)
		_ = container.Provide(ginServer.NewGinServer)
		_ = container.Provide(commonHandler.NewBaseHandler)
		_ = container.Provide(handlers.NewHandlers)

		_ = container.Provide(elasticsearch.NewCoreElkClient)
		_ = container.Provide(gpt.NewGptAdaptorClient)

		_ = container.Provide(repositories.NewResumeRepository)

		_ = container.Provide(services.NewDataProcessingService)

		_ = container.Provide(handlers.NewDataProcessingHandler)

	}

	return container
}

// GetContainer returns an instance of Container.
func GetContainer() *dig.Container {
	return container
}
