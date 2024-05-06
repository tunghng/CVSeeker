package handlers

import (
	"CVSeeker/cmd/CVSeeker/pkg/utils"
	internalDTO "CVSeeker/internal/dtos"
	"CVSeeker/internal/ginLogger"
	"CVSeeker/pkg/api"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// Handlers contains all handlers.
type Handlers struct {
	DataProcessingHandler *DataProcessingHandler
	SearchHandler         *SearchHandler
	ChatbotHandler        *ChatbotHandler
}

// NewHandlersParams contains all dependencies of handlers.
type handlersParams struct {
	dig.In
	DataProcessingHandler *DataProcessingHandler
	SearchHandler         *SearchHandler
	ChatbotHandler        *ChatbotHandler
}

// NewHandlers returns new instance of Handlers.
func NewHandlers(params handlersParams) *Handlers {
	return &Handlers{
		DataProcessingHandler: params.DataProcessingHandler,
		SearchHandler:         params.SearchHandler,
		ChatbotHandler:        params.ChatbotHandler,
	}
}

func GetUserContext(c *gin.Context) *string {
	username := utils.Str2StrPointer(c.GetHeader(api.XForwardUserOpsHeader))
	if username != nil {
		return username
	}
	username = utils.Str2StrPointer(c.GetString(internalDTO.GinContextBasicUsername))
	if username == nil {
		ginLogger.Gin(c).Debugf("Missing username from contexts: %s, %s",
			api.XForwardUserOpsHeader, internalDTO.GinContextBasicUsername)
	}
	return username
}
