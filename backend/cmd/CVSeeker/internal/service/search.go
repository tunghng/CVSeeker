package services

import (
	"CVSeeker/cmd/CVSeeker/internal/cfg"
	"CVSeeker/internal/ginLogger"
	"CVSeeker/internal/meta"
	"CVSeeker/pkg/elasticsearch"
	"CVSeeker/pkg/huggingface"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"net/http"
)

type SearchService interface {
	HybridSearch(c *gin.Context, query string, from, size int, knnBoost float32) (*meta.BasicResponse, error)
	GetDocumentByID(c *gin.Context, documentID string) (*meta.BasicResponse, error)
	DeleteDocumentByID(c *gin.Context, documentID string) (*meta.BasicResponse, error)
}

type searchServiceImpl struct {
	elasticClient elasticsearch.IElasticsearchClient
	hfClient      huggingface.IHuggingFaceClient
}

type SearchServiceArgs struct {
	dig.In
	ElasticClient elasticsearch.IElasticsearchClient
	HfClient      huggingface.IHuggingFaceClient
}

func NewSearchService(args SearchServiceArgs) SearchService {
	return &searchServiceImpl{
		elasticClient: args.ElasticClient,
		hfClient:      args.HfClient,
	}
}

func (_this *searchServiceImpl) HybridSearch(c *gin.Context, query string, from, size int, knnBoost float32) (*meta.BasicResponse, error) {
	textEmbeddingModel := viper.GetString(cfg.HuggingfaceModel)
	indexName := viper.GetString(cfg.ElasticsearchDocumentIndex) // Ensure you configure your index name in viper settings

	// Create the vector representation of text
	vectorEmbedding, err := _this.hfClient.GetTextEmbedding(query, textEmbeddingModel)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to get text embedding: %v", err)
		return nil, err
	}

	// Conduct the hybrid search with pagination
	results, err := _this.elasticClient.HybridSearchWithBoost(c, indexName, query, vectorEmbedding, from, size, knnBoost)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to conduct hybrid search: %v", err)
		return nil, err
	}

	response := &meta.BasicResponse{
		Meta: meta.Meta{
			Code:    http.StatusOK,
			Message: "Search completed successfully",
		},
		Data: results,
	}

	return response, nil
}

func (_this *searchServiceImpl) GetDocumentByID(c *gin.Context, documentID string) (*meta.BasicResponse, error) {
	indexName := viper.GetString(cfg.ElasticsearchDocumentIndex) // Ensure your index name is configured in viper settings

	// Retrieve the document by ID using the Elasticsearch client
	document, err := _this.elasticClient.GetDocumentByID(c, indexName, documentID)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to get document by ID: %v", err)
		return nil, err
	}

	response := &meta.BasicResponse{
		Meta: meta.Meta{
			Code:    http.StatusOK,
			Message: "Document retrieval successful",
		},
		Data: document,
	}

	return response, nil
}

func (_this *searchServiceImpl) DeleteDocumentByID(c *gin.Context, documentID string) (*meta.BasicResponse, error) {
	indexName := viper.GetString(cfg.ElasticsearchDocumentIndex)

	// Delete the document by ID using the Elasticsearch client
	err := _this.elasticClient.DeleteDocumentByID(c, indexName, documentID)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to delete document by ID: %v", err)
		return nil, err
	}

	response := &meta.BasicResponse{
		Meta: meta.Meta{
			Code:    http.StatusOK,
			Message: "Document deletion successful",
		},
		Data: nil, // No data to return for deletion operations
	}

	return response, nil
}
