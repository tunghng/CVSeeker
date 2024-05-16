package handlers

import (
	services "CVSeeker/cmd/CVSeeker/internal/service"
	"CVSeeker/internal/dtos"
	"CVSeeker/internal/errors"
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
// @Param request body dtos.ResumeData true "Resume data including file bytes"
// @Success 200 {object} meta.BasicResponse{data=dtos.ResumeProcessingResult}
// @Failure 400,401,404,500 {object} meta.Error
// @Router /cvseeker/resumes/upload [post]
func (_this *DataProcessingHandler) ProcessDataHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData dtos.ResumeData
		if err := c.ShouldBindJSON(&requestData); err != nil {
			_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInvalidRequest))
			return
		}

		if strings.TrimSpace(requestData.Content) == "" || strings.TrimSpace(requestData.FileBytes) == "" {
			_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInvalidRequest))
			return
		}

		// Process data (example function call, replace with actual processing logic)
		resp, err := _this.dataProcessingService.ProcessData(c, requestData.Content, requestData.FileBytes, requestData.UUID)
		_this.HandleResponse(c, resp, err)
	}
}

// ProcessDataBatchHandler
// @Summary Batch processes resume data
// @Description Processes multiple uploaded resume files and associated metadata as JSON in a single batch.
// @Tags Data Processing
// @Accept json
// @Produce json
// @Param request body dtos.ResumesRequest true "Batch of resume data including file bytes for each"
// @Success 200 {object} meta.BasicResponse{data=[]dtos.ResumeProcessingResult}
// @Failure 400,401,404,500 {object} meta.Error
// @Router /cvseeker/resumes/batch/upload [post]
func (_this *DataProcessingHandler) ProcessDataBatchHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData dtos.ResumesRequest
		if err := c.ShouldBindJSON(&requestData); err != nil {
			_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInvalidRequest))
			return
		}

		if len(requestData.Resumes) == 0 {
			_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInvalidRequest))
			return
		}

		for _, resume := range requestData.Resumes {
			if strings.TrimSpace(resume.Content) == "" || strings.TrimSpace(resume.FileBytes) == "" {
				_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInvalidRequest))
				return
			}
		}

		// Process the batch of resumes
		resp, err := _this.dataProcessingService.ProcessDataBatch(c, requestData.Resumes)
		_this.HandleResponse(c, resp, err)
	}
}

// GetAllUploadsHandler
// @Summary Retrieves all upload records
// @Description Fetches a list of all upload records sorted from the most recent to the oldest
// @Tags Data Processing
// @Accept json
// @Produce json
// @Success 200 {object} meta.BasicResponse{data=[]dtos.UploadDTO}
// @Failure 400,401,404,500 {object} meta.Error
// @Router /cvseeker/resumes/upload [get]
func (_this *DataProcessingHandler) GetAllUploadsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := _this.dataProcessingService.GetAllUploads(c)
		_this.HandleResponse(c, resp, err)
	}
}
