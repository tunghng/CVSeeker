package summarizer

import (
	"CVSeeker/pkg/cfg"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

type ISummarizerAdaptorClient interface {
	AskGPT(prompt, model string) (string, error)
}

type SummarizerAdaptorClient struct {
	Client  *http.Client
	BaseURL string
	ApiKey  string
}

// NewSummarizerAdaptorClient initializes a new client for interacting with GPT models.
func NewSummarizerAdaptorClient(cfgReader *viper.Viper) (ISummarizerAdaptorClient, error) {
	return &SummarizerAdaptorClient{
		Client:  &http.Client{},
		BaseURL: "https://api.openai.com",
		ApiKey:  cfgReader.GetString(cfg.GptApiKey),
	}, nil
}

// addCommonHeaders adds required headers for API authentication.
func (g *SummarizerAdaptorClient) addCommonHeaders(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", g.ApiKey))
}

// AskGPT sends a prompt to the GPT-3.5 API and returns the generated response.
func (g *SummarizerAdaptorClient) AskGPT(prompt, model string) (string, error) {
	endpoint := fmt.Sprintf("%s/v1/chat/completions", g.BaseURL)
	body := map[string]interface{}{
		"model": model,
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
		"temperature": 0.0, // Adjust the temperature if needed
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("could not encode request body: %v", err)
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", fmt.Errorf("could not create request: %v", err)
	}

	g.addCommonHeaders(req)

	resp, err := g.Client.Do(req)
	if err != nil {
		return "", fmt.Errorf("could not send request to GPT API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := json.Marshal(body) // assuming error handling omitted for brevity
		return "", fmt.Errorf("GPT API returned non-OK status code: %d, message: %s", resp.StatusCode, string(bodyBytes))
	}

	var response struct {
		Choices []struct {
			Message struct {
				Role    string `json:"role"`
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", fmt.Errorf("could not decode response body: %v", err)
	}

	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response found in the API return")
}
