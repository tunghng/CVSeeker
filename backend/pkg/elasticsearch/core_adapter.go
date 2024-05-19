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
	AddDocument(ctx context.Context, indexName string, document interface{}) (string, error)
	KeywordSearch(ctx context.Context, indexName string, query string) ([]ResumeSummaryDTO, error)
	VectorSearch(ctx context.Context, indexName string, vector []float32) ([]ResumeSummaryDTO, error)
	DeleteDocumentByID(ctx context.Context, indexName, documentID string) error
	HybridSearchWithBoost(ctx context.Context, indexName, query string, queryVector []float32, from, size int, knnBoost float32) ([]ResumeSummaryDTO, error)
	GetDocumentByID(ctx context.Context, indexName, documentId string) (*ResumeSummaryDTO, error)
	FetchDocumentsByIDs(ctx context.Context, indexName string, documentIDs []string) ([]ResumeSummaryDTO, error)
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
		Transport: customTransport,
		Logger:    &elastictransport.ColorLogger{Output: os.Stdout, EnableRequestBody: true, EnableResponseBody: true},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create elasticsearch client: %w", err)
	}

	return &ElasticsearchClient{client: es}, nil
}

// AddDocument adds a new document to the specified index
func (ec *ElasticsearchClient) AddDocument(ctx context.Context, indexName string, document interface{}) (string, error) {
	docJSON, err := json.Marshal(document)
	if err != nil {
		return "", fmt.Errorf("error marshaling document: %w", err)
	}

	// Prepare the request with the specified index, document body, and make it refresh immediately
	req := esapi.IndexRequest{
		Index:   indexName,
		Body:    bytes.NewReader(docJSON),
		Refresh: "true", // or use esapi.RefreshTrue if available
	}

	// Perform the request with the given context
	res, err := req.Do(ctx, ec.client)
	if err != nil {
		return "", fmt.Errorf("error indexing document: %w", err)
	}
	defer res.Body.Close()

	// Read the response body
	if res.IsError() {
		return "", fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}

	// Parse the response body to extract the ID
	var result struct {
		ID string `json:"_id"`
	}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error parsing response body: %w", err)
	}

	return result.ID, nil
}

// GetDocumentByID retrieves a document by its ID from a specific index and converts it to an ResumeSummaryDTO.
func (ec *ElasticsearchClient) GetDocumentByID(ctx context.Context, indexName string, documentID string) (*ResumeSummaryDTO, error) {
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

func (ec *ElasticsearchClient) FetchDocumentsByIDs(ctx context.Context, indexName string, documentIDs []string) ([]ResumeSummaryDTO, error) {
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

	// Convert the results to ResumeSummaryDTO
	response := make([]ResumeSummaryDTO, 0, len(mgetResp.Docs))
	for _, doc := range mgetResp.Docs {
		if doc.Found {
			var resume ResumeSummaryDTO

			// Handle content and additional fields
			contentData, ok := doc.Source["content"].(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("content field missing or not correctly formatted as a JSON object")
			}

			jsonData, err := json.Marshal(contentData)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal content data: %w", err)
			}

			if err := json.Unmarshal(jsonData, &resume); err != nil {
				return nil, fmt.Errorf("error unmarshaling resume content: %w", err)
			}

			// Assign the document ID
			resume.Id = doc.ID

			response = append(response, resume)
		}
	}

	return response, nil
}

func (ec *ElasticsearchClient) DeleteDocumentByID(ctx context.Context, indexName, documentID string) error {
	// Create the Delete request to Elasticsearch
	req := esapi.DeleteRequest{
		Index:      indexName,
		DocumentID: documentID,
	}

	// Perform the request with the provided context
	res, err := req.Do(ctx, ec.client)
	if err != nil {
		return fmt.Errorf("error deleting document: %w", err)
	}
	defer res.Body.Close() // Ensure body is closed after the operation

	if res.IsError() {
		return fmt.Errorf("error response from Elasticsearch while deleting document: %s", res.String())
	}

	return nil
}

func (ec *ElasticsearchClient) KeywordSearch(ctx context.Context, indexName string, query string) ([]ResumeSummaryDTO, error) {
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

func (ec *ElasticsearchClient) VectorSearch(ctx context.Context, indexName string, vector []float32) ([]ResumeSummaryDTO, error) {
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
func (ec *ElasticsearchClient) HybridSearchWithBoost(ctx context.Context, indexName, query string, queryVector []float32, from, size int, knnBoost float32) ([]ResumeSummaryDTO, error) {
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
			K:             60,
			NumCandidates: 200, // Adj	ust the number of candidates as needed
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

func ConvertHitsToElasticResponses(hits []types.Hit) ([]ResumeSummaryDTO, error) {
	var resumes []ResumeSummaryDTO
	for _, hit := range hits {
		resume, err := ConvertHitToElasticResponse(&hit)
		if err != nil {
			return nil, err
		}
		resumes = append(resumes, *resume)
	}
	return resumes, nil
}

func ConvertHitToElasticResponse(hit *types.Hit) (*ResumeSummaryDTO, error) {
	var source map[string]interface{}
	if err := json.Unmarshal(hit.Source_, &source); err != nil {
		return nil, fmt.Errorf("an error occurred while unmarshaling hit: %w", err)
	}

	// Extract the ID from the hit and assign it to the DTO
	id := hit.Id_

	score := float64(hit.Score_)

	// Check if the content is a JSON object and handle it directly
	contentData, ok := source["content"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("content field missing or not correctly formatted as a JSON object")
	}

	// Marshal the contentData back to JSON string to unmarshal into DTO
	jsonData, err := json.Marshal(contentData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal content data: %w", err)
	}

	var resume ResumeSummaryDTO
	if err := json.Unmarshal(jsonData, &resume); err != nil {
		return nil, fmt.Errorf("an error occurred while unmarshaling resume content: %w", err)
	}

	resume.Id = id
	resume.Point = score

	return &resume, nil
}
