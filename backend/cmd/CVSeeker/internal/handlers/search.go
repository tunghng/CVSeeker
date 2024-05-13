package handlers

import (
	services "CVSeeker/cmd/CVSeeker/internal/service"
	"CVSeeker/internal/dtos"
	"CVSeeker/internal/errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"strconv"
)

type SearchHandler struct {
	BaseHandler
	searchService services.SearchService
}

type SearchHandlerParams struct {
	dig.In
	BaseHandler   BaseHandler
	SearchService services.SearchService
}

func NewSearchHandler(params SearchHandlerParams) *SearchHandler {
	return &SearchHandler{
		searchService: params.SearchService,
		BaseHandler:   params.BaseHandler,
	}
}

// HybridSearch
// @Summary Perform hybridsearch on elasticsearch
// @Description Executes a search combining keyword and vector-based queries with customizable boosting on the vector component.
// @Tags Search
// @Accept json
// @Produce json
// @Param body body dtos.QueryRequest true "Message content"
// @Param knnBoost query float32 false "Boost factor for the KNN component" default(0.5)
// @Param from query int false "Start index for search results" default(0)
// @Param size query int false "Number of search results to return" default(10)
// @Success 200 {object} meta.BasicResponse
// @Failure 400,401,404,500 {object} meta.Error
// @Router /cvseeker/resumes/search [GET]
func (_this *SearchHandler) HybridSearch() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request dtos.QueryRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInvalidRequest))
			return
		}

		knnBoost, err := strconv.ParseFloat(c.DefaultQuery("knnBoost", "0.5"), 32)
		if err != nil {
			_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInvalidRequest))
			return
		}

		from, err := strconv.Atoi(c.DefaultQuery("from", "0"))
		if err != nil {
			_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInvalidRequest))
			return
		}

		size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
		if err != nil {
			_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInvalidRequest))
			return
		}

		resp, err := _this.searchService.HybridSearch(c, request.Content, from, size, float32(knnBoost))
		if err != nil {
			_this.HandleResponse(c, nil, err)
			return
		}
		_this.HandleResponse(c, resp, nil)
	}
}

// GetDocumentByID
// @Summary Get Document By Id
// @Description Retrieves a document by its ID from the Elasticsearch index.
// @Tags Search
// @Accept json
// @Produce json
// @Param id path string true "Document ID"
// @Success 200 {object} meta.BasicResponse
// @Failure 400,401,404,500 {object} meta.Error
// @Router /cvseeker/resumes/{id} [GET]
func (_this *SearchHandler) GetDocumentByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get document ID from query parameters or path parameters
		documentID := c.Param("id") // Assuming the ID is passed as a URL parameter
		if documentID == "" {
			_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInvalidRequest))
			return
		}

		response, err := _this.searchService.GetDocumentByID(c, documentID)
		_this.HandleResponse(c, response, err)
	}
}

// DeleteDocumentByID
// @Summary Delete Document By Id
// @Description Deletes a document by its ID from the Elasticsearch index.
// @Tags Search
// @Accept json
// @Produce json
// @Param id path string true "Document ID"
// @Success 200 {object} meta.BasicResponse "Document deletion successful"
// @Failure 400,401,404,500 {object} meta.Error
// @Router /cvseeker/resumes/{id} [DELETE]
func (_this *SearchHandler) DeleteDocumentByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get document ID from path parameters
		documentID := c.Param("id")
		if documentID == "" {
			_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInvalidRequest))
			return
		}

		// Perform deletion through the search service
		resp, err := _this.searchService.DeleteDocumentByID(c, documentID)
		_this.HandleResponse(c, resp, err)
	}
}
