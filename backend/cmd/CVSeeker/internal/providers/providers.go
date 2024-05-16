package providers

import (
	"CVSeeker/cmd/CVSeeker/internal/handlers"
	services "CVSeeker/cmd/CVSeeker/internal/service"
	"CVSeeker/internal/errors"
	"CVSeeker/internal/ginServer"
	"CVSeeker/internal/repositories"
	"CVSeeker/pkg/aws"
	"CVSeeker/pkg/cfg"
	"CVSeeker/pkg/elasticsearch"
	"CVSeeker/pkg/gpt"
	"CVSeeker/pkg/huggingface"
	"CVSeeker/pkg/logger"
	"CVSeeker/pkg/summarizer"
	"CVSeeker/pkg/websocket"
	"go.uber.org/dig"
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
		_ = container.Provide(newGinEngine)
		_ = container.Provide(setupRouter)
		_ = container.Provide(newServerConfig)
		_ = container.Provide(newErrorParserConfig)
		_ = container.Provide(newMySQLConnection, dig.Name("talentAcquisitionDB"))

		_ = container.Provide(logger.NewLogger)
		_ = container.Provide(errors.NewErrorParser)
		_ = container.Provide(ginServer.NewGinServer)
		_ = container.Provide(handlers.NewBaseHandler)
		_ = container.Provide(handlers.NewHandlers)

		_ = container.Provide(elasticsearch.NewElasticsearchClient)
		_ = container.Provide(summarizer.NewSummarizerAdaptorClient)
		_ = container.Provide(huggingface.NewHuggingFaceClient)
		_ = container.Provide(aws.NewS3Client)
		_ = container.Provide(gpt.NewGptAdaptorClient)
		_ = container.Provide(websocket.NewWebSocketClient)

		_ = container.Provide(repositories.NewResumeRepository)
		_ = container.Provide(repositories.NewThreadResumeRepository)
		_ = container.Provide(repositories.NewThreadRepository)
		_ = container.Provide(repositories.NewUploadRepository)

		_ = container.Provide(services.NewDataProcessingService)
		_ = container.Provide(services.NewSearchService)
		_ = container.Provide(services.NewChatbotService)

		_ = container.Provide(handlers.NewDataProcessingHandler)
		_ = container.Provide(handlers.NewSearchHandler)
		_ = container.Provide(handlers.NewChatbotHandler)
	}

	return container
}

// GetContainer returns an instance of Container.
func GetContainer() *dig.Container {
	return container
}
