package handlers

import (
	services "CVSeeker/cmd/CVSeeker/internal/service"
	"CVSeeker/internal/handlers"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"io/ioutil"
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
func (_this *DataProcessingHandler) ProcessDataHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fullText := c.PostForm("fullText")
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			_this.RespondError(c, err)
			return
		}

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			_this.RespondError(c, err)
			return
		}

		resp, err := _this.dataProcessingService.ProcessData(c, fullText, fileBytes)
		_this.HandleResponse(c, resp, err)
	}
}
