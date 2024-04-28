package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"grabber-match/cmd/CVSeeker/internal/service"
	"grabber-match/internal/errors"
	"grabber-match/internal/handlers"
)

type DataProcessingHandler struct {
	handlers.BaseHandler
	dataProcessingService services.IDataProcessingService
}

type DataProcessingHandlerParams struct {
	dig.In
	BaseHandler           handlers.BaseHandler
	DataProcessingService services.IDataProcessingService
}

func NewDataProcessingHandler(params DataProcessingHandlerParams) *DataProcessingHandler {
	return &DataProcessingHandler{
		BaseHandler:           params.BaseHandler,
		dataProcessingService: params.DataProcessingService,
	}
}

// / HandleSummarizeResume is the Gin handler function to summarize resumes.
func (_this *DataProcessingHandler) HandleSummarizeResume() gin.HandlerFunc {
	return func(c *gin.Context) {
		fullText := c.Query("fullText")

		if fullText == "" {
			_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInternalServer))
			return
		}

		resp, err := _this.dataProcessingService.SummarizeResume(fullText)
		_this.HandleResponse(c, resp, err)
	}
}
