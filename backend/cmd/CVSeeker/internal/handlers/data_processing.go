package handlers

import (
	services "CVSeeker/cmd/CVSeeker/internal/service"
	"CVSeeker/internal/errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"io"
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
// @Description Processes uploaded resume files and associated metadata
// @Tags Data Processing
// @Accept multipart/form-data
// @Produce json
// @Param fullText query string true "Full text of the resume"
// @Param file formData file true "Upload file"
// @Success 200 {object} meta.BasicResponse
// @Failure 400,401,404,500 {object} meta.Error
// @Router /cvseeker/resumes/update [post]
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
