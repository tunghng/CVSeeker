package elasticsearch

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/spf13/viper"
	"net/http"
	"os"

	"CVSeeker/pkg/cfg"
)

type IElasticsearchClient interface {
	AddDocument(ctx context.Context, indexName string, documentId string, document interface{}) error
	KeywordSearch(ctx context.Context, indexName string, term string) ([]ElasticResponse, error)
	VectorSearch(ctx context.Context, indexName string, vector []float32) ([]ElasticResponse, error)
	HybridSearchWithBoost(ctx context.Context, indexName, term string, queryVector []float32, knnBoost float32, numResults int) ([]ElasticResponse, error)
}

type ElasticsearchClient struct {
	client *elasticsearch.TypedClient
}

func NewElasticsearchClient(cfgReader *viper.Viper) (IElasticsearchClient, error) {
	url := cfgReader.GetString(cfg.ElasticsearchUrl)
	username := cfgReader.GetString(cfg.ElasticsearchUserName)
	password := cfgReader.GetString(cfg.ElasticsearchPassword)

	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // or configure as needed
	}

	es, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{url},
		Username:  username,
		Password:  password,
		Transport: customTransport, // Using custom transport
		Logger:    &elastictransport.ColorLogger{Output: os.Stdout, EnableRequestBody: true, EnableResponseBody: true},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create elasticsearch client: %w", err)
	}

	return &ElasticsearchClient{client: es}, nil
}

// AddDocument adds a new document to the specified index
func (ec *ElasticsearchClient) AddDocument(ctx context.Context, indexName string, documentId string, document interface{}) error {
	docJSON, err := json.Marshal(document)
	if err != nil {
		return fmt.Errorf("error marshaling document: %w", err)
	}

	// Prepare the request with the specified index, document ID and document body
	req := esapi.IndexRequest{
		Index:      indexName,
		DocumentID: documentId,
		Body:       bytes.NewReader(docJSON),
		Refresh:    "true", // or use esapi.RefreshTrue if available
	}

	// Perform the request with the given context
	res, err := req.Do(ctx, ec.client)
	if err != nil {
		return fmt.Errorf("error indexing document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}

	return nil
}

func (ec *ElasticsearchClient) KeywordSearch(ctx context.Context, indexName string, term string) ([]ElasticResponse, error) {
	res, err := ec.client.Search().
		Index(indexName).
		Query(&types.Query{
			Match: map[string]types.MatchQuery{
				"content": {Query: term},
			},
		}).
		Do(ctx)

	if err != nil {
		return nil, fmt.Errorf("keyword search failed: %w", err)
	}

	return ConvertHitsToElasticResponses(res.Hits.Hits)
}

func (ec *ElasticsearchClient) VectorSearch(ctx context.Context, indexName string, vector []float32) ([]ElasticResponse, error) {
	res, err := ec.client.Search().
		Index(indexName).
		Knn(types.KnnQuery{
			Field:       "embedding",
			QueryVector: vector,
			K:           10,
		}).
		Do(ctx)

	if err != nil {
		return nil, fmt.Errorf("vector search failed: %w", err)
	}

	return ConvertHitsToElasticResponses(res.Hits.Hits)
}

func (ec *ElasticsearchClient) HybridSearchWithBoost(ctx context.Context, indexName, term string, queryVector []float32, knnBoost float32, numResults int) ([]ElasticResponse, error) {
	queryBoost := 1.0 - knnBoost

	// Generate a query vector for the term, replace this with your actual model vector generation
	res, err := ec.client.Search().
		Index(indexName).
		Size(numResults).
		Knn(types.KnnQuery{
			Field:         "embedding", // Ensure this field matches your schema
			QueryVector:   queryVector,
			Boost:         &knnBoost,
			K:             10,
			NumCandidates: 100, // Adjust the number of candidates as needed
		}).
		Query(&types.Query{
			Match: map[string]types.MatchQuery{
				"content": {
					Query: term,
					Boost: &queryBoost,
				},
			},
		}).
		Do(ctx)

	if err != nil {
		return nil, fmt.Errorf("hybrid search failed: %w", err)
	}

	return ConvertHitsToElasticResponses(res.Hits.Hits)
}

func ConvertHitsToElasticResponses(hits []types.Hit) ([]ElasticResponse, error) {
	var responses []ElasticResponse
	for _, hit := range hits {
		var source map[string]interface{} // Use interface{} to accept any data type
		if err := json.Unmarshal(hit.Source_, &source); err != nil {
			return nil, fmt.Errorf("an error occurred while unmarshaling hit: %w", err)
		}

		var content, url string
		// Check if content is a string or a slice and assign accordingly
		if contentVal, ok := source["content"].([]interface{}); ok && len(contentVal) > 0 {
			content, _ = contentVal[0].(string) // Safely assert to string
		} else if contentStr, ok := source["content"].(string); ok {
			content = contentStr
		}

		// Check if url is a string or a slice and assign accordingly
		if urlVal, ok := source["url"].([]interface{}); ok && len(urlVal) > 0 {
			url, _ = urlVal[0].(string) // Safely assert to string
		} else if urlStr, ok := source["url"].(string); ok {
			url = urlStr
		}

		response := ElasticResponse{
			ID:      hit.Id_,
			Content: content,
			URL:     url,
		}
		responses = append(responses, response)
	}
	return responses, nil
}
