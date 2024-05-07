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
	KeywordSearch(ctx context.Context, indexName string, query string) ([]ElasticResponse, error)
	VectorSearch(ctx context.Context, indexName string, vector []float32) ([]ElasticResponse, error)
	HybridSearchWithBoost(ctx context.Context, indexName, query string, queryVector []float32, from, size int, knnBoost float32) ([]ElasticResponse, error)
	GetDocumentByID(ctx context.Context, indexName, documentId string) (*ElasticResponse, error)
	FetchDocumentsByIDs(ctx context.Context, indexName string, documentIDs []string) ([]ElasticResponse, error)
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

// GetDocumentByID retrieves a document by its ID from a specific index and converts it to an ElasticResponse.
func (ec *ElasticsearchClient) GetDocumentByID(ctx context.Context, indexName string, documentID string) (*ElasticResponse, error) {
	// Create the Get request to Elasticsearch
	req := esapi.GetRequest{
		Index:      indexName,
		DocumentID: documentID,
	}

	// Perform the request with the provided context
	res, err := req.Do(ctx, ec.client)
	if err != nil {
		return nil, fmt.Errorf("error retrieving document: %w", err)
	}
	defer res.Body.Close() // Ensure body is closed after the operation

	// Check if the request was not successful
	if !res.IsError() {
		var hit types.Hit
		if err := json.NewDecoder(res.Body).Decode(&hit); err != nil {
			return nil, fmt.Errorf("error decoding response body: %w", err)
		}
		// Convert the hit to an ElasticResponse
		return ConvertHitToElasticResponse(&hit)
	} else {
		return nil, fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}
}

func (ec *ElasticsearchClient) FetchDocumentsByIDs(ctx context.Context, indexName string, documentIDs []string) ([]ElasticResponse, error) {
	// Construct the request body for the multi-get API
	docs := make([]map[string]interface{}, len(documentIDs))
	for i, id := range documentIDs {
		docs[i] = map[string]interface{}{
			"_id": id,
		}
	}
	requestBody, err := json.Marshal(map[string]interface{}{"docs": docs})
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}

	// Perform the multi-get request
	req := esapi.MgetRequest{
		Index: indexName,
		Body:  bytes.NewReader(requestBody),
	}

	res, err := req.Do(ctx, ec.client)
	if err != nil {
		return nil, fmt.Errorf("error performing mget request: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}

	// Decode the response
	var mgetResp struct {
		Docs []struct {
			Found  bool                   `json:"found"`
			Source map[string]interface{} `json:"_source"`
			ID     string                 `json:"_id"`
		} `json:"docs"`
	}
	if err := json.NewDecoder(res.Body).Decode(&mgetResp); err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}

	// Convert the results to ElasticResponse
	responses := make([]ElasticResponse, 0, len(mgetResp.Docs))
	for _, doc := range mgetResp.Docs {
		if doc.Found {
			response := ElasticResponse{
				ID:      doc.ID,
				Content: fmt.Sprintf("%v", doc.Source["content"]),
				URL:     fmt.Sprintf("%v", doc.Source["url"]),
			}
			responses = append(responses, response)
		}
	}

	return responses, nil
}

func (ec *ElasticsearchClient) KeywordSearch(ctx context.Context, indexName string, query string) ([]ElasticResponse, error) {
	res, err := ec.client.Search().
		Index(indexName).
		Query(&types.Query{
			Match: map[string]types.MatchQuery{
				"content": {Query: query},
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

// HybridSearchWithBoost perform search combining both semantic and lexiacal search
func (ec *ElasticsearchClient) HybridSearchWithBoost(ctx context.Context, indexName, query string, queryVector []float32, from, size int, knnBoost float32) ([]ElasticResponse, error) {
	queryBoost := 1.0 - knnBoost

	// Generate a query vector for the term, replace this with your actual model vector generation
	res, err := ec.client.Search().
		Index(indexName).
		From(from).
		Size(size).
		Knn(types.KnnQuery{
			Field:         "embedding", // Ensure this field matches your schema
			QueryVector:   queryVector,
			Boost:         &knnBoost,
			K:             100,
			NumCandidates: 100, // Adj	ust the number of candidates as needed
		}).
		Query(&types.Query{
			Match: map[string]types.MatchQuery{
				"content": {
					Query: query,
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
		response, err := ConvertHitToElasticResponse(&hit)
		if err != nil {
			return nil, err
		}
		responses = append(responses, *response)
	}
	return responses, nil
}

// ConvertHitToElasticResponse converts a single Elasticsearch hit to an ElasticResponse.
func ConvertHitToElasticResponse(hit *types.Hit) (*ElasticResponse, error) {
	var source map[string]interface{}
	if err := json.Unmarshal(hit.Source_, &source); err != nil {
		return nil, fmt.Errorf("an error occurred while unmarshaling hit: %w", err)
	}

	var content, url string
	// Handle content field based on its data type in the source
	if contentVal, ok := source["content"].([]interface{}); ok && len(contentVal) > 0 {
		content, _ = contentVal[0].(string) // Safely assert to string
	} else if contentStr, ok := source["content"].(string); ok {
		content = contentStr
	}

	// Handle url field based on its data type in the source
	if urlVal, ok := source["url"].([]interface{}); ok && len(urlVal) > 0 {
		url, _ = urlVal[0].(string) // Safely assert to string
	} else if urlStr, ok := source["url"].(string); ok {
		url = urlStr
	}

	response := &ElasticResponse{
		ID:      hit.Id_,
		Content: content,
		URL:     url,
	}

	return response, nil
}
