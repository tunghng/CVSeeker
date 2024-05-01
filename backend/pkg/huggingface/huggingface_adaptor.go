package huggingface

import (
	"CVSeeker/pkg/cfg"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

type IHuggingFaceClient interface {
	GetTextEmbedding(term string, model string) ([]float32, error)
}

type HuggingFaceClient struct {
	httpClient *http.Client
	apiKey     string
}

func NewHuggingFaceClient(cfgReader *viper.Viper) (IHuggingFaceClient, error) {
	return &HuggingFaceClient{
		httpClient: &http.Client{},
		apiKey:     cfgReader.GetString(cfg.HuggingfaceApiKey),
	}, nil
}

func (hc *HuggingFaceClient) GetTextEmbedding(term string, model string) ([]float32, error) {
	posturl := fmt.Sprintf("https://api-inference.huggingface.co/pipeline/feature-extraction/%s", model)

	// Prepare the request body with JSON content
	body := []byte(fmt.Sprintf(`{"inputs": "%s", "options": {"wait_for_model": true}}`, term))

	// Create a HTTP post request
	req, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+hc.apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := hc.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read and decode the response body
	var embeddings []float32
	if err := json.NewDecoder(resp.Body).Decode(&embeddings); err != nil {
		return nil, err
	}

	return embeddings, nil
}
