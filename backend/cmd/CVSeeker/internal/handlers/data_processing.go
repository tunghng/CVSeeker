package handlers

import (
	services "CVSeeker/cmd/CVSeeker/internal/service"
	"CVSeeker/internal/dtos"
	"CVSeeker/internal/errors"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"strings"
)

type DataProcessingHandler struct {
	BaseHandler
	dataProcessingService services.IDataProcessingService
}

type DataProcessingHandlerParams struct {
	dig.In
	BaseHandler           BaseHandler
	DataProcessingService services.IDataProcessingService
}

func NewDataProcessingHandler(params DataProcessingHandlerParams) *DataProcessingHandler {
	return &DataProcessingHandler{
		BaseHandler:           params.BaseHandler,
		dataProcessingService: params.DataProcessingService,
	}
}

// ProcessDataHandler
// @Summary Processes resume data
// @Description Processes uploaded resume files and associated metadata as JSON
// @Tags Data Processing
// @Accept json
// @Produce json
// @Param request body dtos.ResumeRequest true "Resume data including file bytes"
// @Success 200 {object} meta.BasicResponse
// @Failure 400,401,404,500 {object} meta.Error
// @Router /cvseeker/resumes/upload [post]
func (_this *DataProcessingHandler) ProcessDataHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData dtos.ResumeRequest
		if err := c.ShouldBindJSON(&requestData); err != nil {
			_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInvalidRequest))
			return
		}

		if strings.TrimSpace(requestData.Content) == "" || strings.TrimSpace(requestData.FileBytes) == "" {
			_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInvalidRequest))
			return
		}

		// Decode base64 file bytes
		fileBytes, err := base64.StdEncoding.DecodeString(requestData.FileBytes)
		if err != nil {
			_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInvalidRequest))
			return
		}

		// Process data (example function call, replace with actual processing logic)
		resp, err := _this.dataProcessingService.ProcessData(c, requestData.Content, fileBytes)
		_this.HandleResponse(c, resp, err)
	}
}
