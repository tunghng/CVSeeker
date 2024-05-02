package handlers

import (
	services "CVSeeker/cmd/CVSeeker/internal/service"
	"CVSeeker/internal/errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"strconv"
	"strings"
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

// HybridSearch handles search requests and interacts with the SearchService.
func (_this *SearchHandler) HybridSearch() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve search term and other parameters from the query string
		query := strings.TrimSpace(c.Query("query"))
		if query == "" {
			_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInvalidRequest))
			return
		}
		// Get knnBoost with a default value
		knnBoost, _ := strconv.ParseFloat(c.DefaultQuery("knnBoost", "0.5"), 32)
		// Get numResults with a default value
		numResults, _ := strconv.Atoi(c.DefaultQuery("numResults", "10"))

		// Call the search service
		resp, err := _this.searchService.HybridSearch(c, query, float32(knnBoost), numResults)
		_this.HandleResponse(c, resp, err)
	}
}

// GetDocumentByID handles requests to retrieve a document by its ID.
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
