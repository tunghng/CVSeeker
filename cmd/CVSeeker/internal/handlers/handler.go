package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"grabber-match/cmd/CVSeeker/pkg/utils"
	internalDTO "grabber-match/internal/dtos"
	"grabber-match/internal/ginLogger"
	"grabber-match/pkg/api"
)

// Handlers contains all handlers.
type Handlers struct {
	DataProcessingHandler *DataProcessingHandler
}

// NewHandlersParams contains all dependencies of handlers.
type handlersParams struct {
	dig.In
	DataProcessingHandler *DataProcessingHandler
}

// NewHandlers returns new instance of Handlers.
func NewHandlers(params handlersParams) *Handlers {
	return &Handlers{
		DataProcessingHandler: params.DataProcessingHandler,
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
