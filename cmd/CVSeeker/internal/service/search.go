package services

import (
	"CVSeeker/internal/ginLogger"
	"CVSeeker/internal/meta"
	"CVSeeker/pkg/elasticsearch"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/dig"
)

type SearchService interface {
	HybridSearch(c *gin.Context, term string, knnBoost float32, numResults int) (*meta.BasicResponse, error)
}

type searchServiceImpl struct {
	elasticClient elasticsearch.ElasticsearchClient
}

type SearchServiceArgs struct {
	dig.In
	ElasticClient elasticsearch.ElasticsearchClient
}

func NewSearchService(args SearchServiceArgs) SearchService {
	return &searchServiceImpl{
		elasticClient: args.ElasticClient,
	}
}

func (s *searchServiceImpl) HybridSearch(c *gin.Context, term string, knnBoost float32, numResults int) (*meta.BasicResponse, error) {
	indexName := viper.GetString("elasticsearch.indexName") // Ensure you configure your index name in viper settings

	// Generate a query vector here based on your model needs; placeholder for now
	queryVector := generateQueryVector(term) // This function should generate a query vector for the term

	// Conduct the hybrid search
	results, err := s.elasticClient.HybridSearchWithBoost(context.Background(), indexName, term, queryVector, knnBoost, numResults)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to conduct hybrid search: %v", err)
		return nil, err
	}

	response := &meta.BasicResponse{
		Meta: meta.Meta{
			Code:    200,
			Message: "Search completed successfully",
		},
		Data: results,
	}

	return response, nil
}

func generateQueryVector(term string) []float32 {
	// Stub for generating a query vector based on input term
	return []float32{} // This should be replaced with actual logic to convert term into a vector
}
