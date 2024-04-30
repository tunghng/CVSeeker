package handlers

import (
	services "CVSeeker/cmd/CVSeeker/internal/service"
	"CVSeeker/internal/handlers"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"strconv"
	"strings"
)

type SearchHandler struct {
	handlers.BaseHandler
	searchService services.SearchService
}

type SearchHandlerParams struct {
	dig.In
	BaseHandler   handlers.BaseHandler
	SearchService services.SearchService
}

func NewSearchHandler(params SearchHandlerParams) *SearchHandler {
	return &SearchHandler{
		searchService: params.SearchService,
		BaseHandler:   params.BaseHandler,
	}
}

// HybridSearchHandler handles search requests and interacts with the SearchService.
func (_this *SearchHandler) HybridSearchHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve search term and other parameters from the query string
		term := strings.TrimSpace(c.Query("term"))
		// Get knnBoost with a default value
		knnBoost, _ := strconv.ParseFloat(c.DefaultQuery("knnBoost", "0.5"), 32)
		// Get numResults with a default value
		numResults, _ := strconv.Atoi(c.DefaultQuery("numResults", "10"))

		// Call the search service
		resp, err := _this.searchService.HybridSearch(c, term, float32(knnBoost), numResults)
		_this.HandleResponse(c, resp, err)
	}
}
