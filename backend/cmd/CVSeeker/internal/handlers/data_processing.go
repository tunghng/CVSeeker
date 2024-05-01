package handlers

import (
	services "CVSeeker/cmd/CVSeeker/internal/service"
	"CVSeeker/internal/errors"
	"CVSeeker/internal/handlers"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"io"
	"strings"
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

// ProcessDataHandler is the Gin handler processing the resume
func (_this *DataProcessingHandler) ProcessDataHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fullText := strings.TrimSpace(c.Query("fullText"))
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInternalServer))
			return
		}

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			_this.RespondError(c, err)
			return
		}

		resp, err := _this.dataProcessingService.ProcessData(c, fullText, fileBytes)
		_this.HandleResponse(c, resp, err)
	}
}
