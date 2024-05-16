package gpt

import (
	"CVSeeker/pkg/cfg"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type IGptAdaptorClient interface {
	CreateAssistant(request AssistantRequest) (*AssistantResponse, error)
	CreateThread(request CreateThreadRequest) (*ThreadResponse, error)
	DeleteThread(threadID string) (*DeleteThreadResponse, error)
	ListMessages(threadID string, limit int, order, after, before string) (*ListMessagesResponse, error)
	GetRunDetails(threadID, runID string) (*RunResponse, error)
	CreateRunAndStreamResponse(threadID string, request CreateRunRequest) (<-chan string, error)

	CreateMessage(threadID string, request CreateMessageRequest) (*MessageResponse, error)
	WaitForRunCompletion(threadID, runID string) (*RunResponse, error)
}

type gptAdaptorClient struct {
	Client *http.Client
	ApiKey string
}

func NewGptAdaptorClient(cfgReader *viper.Viper) (IGptAdaptorClient, error) {
	return &gptAdaptorClient{
		Client: &http.Client{},
		ApiKey: cfgReader.GetString(cfg.GptApiKey),
	}, nil
}

func (g *gptAdaptorClient) addCommonHeaders(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+g.ApiKey)
	req.Header.Add("OpenAI-Beta", OpenaiAssistantsV1)
}

func (g *gptAdaptorClient) CreateAssistant(request AssistantRequest) (*AssistantResponse, error) {
	url := AssistantEndpoint

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	g.addCommonHeaders(req)

	resp, err := g.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		var errMsg string
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			errMsg = "Failed to read response body"
		} else {
			errMsg = string(bodyBytes)
		}
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, errMsg)
	}
	var response AssistantResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (g *gptAdaptorClient) CreateThread(request CreateThreadRequest) (*ThreadResponse, error) {
	url := ThreadEndpoint

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	g.addCommonHeaders(req)

	resp, err := g.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		var errMsg string
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			errMsg = "Failed to read response body"
		} else {
			errMsg = string(bodyBytes)
		}
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, errMsg)
	}
	// Xử lý response
	var response ThreadResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (g *gptAdaptorClient) DeleteThread(threadID string) (*DeleteThreadResponse, error) {
	url := fmt.Sprintf("%v/%v", ThreadEndpoint, threadID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	g.addCommonHeaders(req)

	resp, err := g.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		var errMsg string
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			errMsg = "Failed to read response body"
		} else {
			errMsg = string(bodyBytes)
		}
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, errMsg)
	}
	// Xử lý response
	var response DeleteThreadResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (g *gptAdaptorClient) CreateMessage(threadID string, request CreateMessageRequest) (*MessageResponse, error) {
	url := fmt.Sprintf("%v/%v/messages", ThreadEndpoint, threadID)

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	g.addCommonHeaders(req)

	resp, err := g.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		var errMsg string
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			errMsg = "Failed to read response body"
		} else {
			errMsg = string(bodyBytes)
		}
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, errMsg)
	}

	var response MessageResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (g *gptAdaptorClient) ListMessages(threadID string, limit int, order, after, before string) (*ListMessagesResponse, error) {
	urls := fmt.Sprintf("%v/%v/messages", ThreadEndpoint, threadID)

	// Xây dựng các tham số truy vấn
	queryParams := url.Values{}
	if limit > 0 {
		queryParams.Add("limit", strconv.Itoa(limit))
	}
	if order != "" {
		queryParams.Add("order", order)
	}
	if after != "" {
		queryParams.Add("after", after)
	}
	if before != "" {
		queryParams.Add("before", before)
	}
	urls += "?" + queryParams.Encode()
	req, err := http.NewRequest("GET", urls, nil)
	if err != nil {
		return nil, err
	}
	g.addCommonHeaders(req)

	resp, err := g.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		var errMsg string
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			errMsg = "Failed to read response body"
		} else {
			errMsg = string(bodyBytes)
		}
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, errMsg)
	}
	// Xử lý response
	var response ListMessagesResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (g *gptAdaptorClient) CreateRunAndStreamResponse(threadID string, request CreateRunRequest) (<-chan string, error) {
	url := fmt.Sprintf("%v/%v/runs", ThreadEndpoint, threadID)

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(requestBody)))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	g.addCommonHeaders(req)

	resp, err := g.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(bodyBytes))
	}

	valueChannel := make(chan string)

	go func() {
		defer close(valueChannel)
		defer resp.Body.Close()

		scanner := bufio.NewScanner(resp.Body)
		var currentEvent string
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "event:") {
				currentEvent = strings.TrimSpace(line[6:])
			} else if strings.HasPrefix(line, "data:") {
				data := strings.TrimSpace(line[5:])
				if currentEvent == "thread.message.delta" {
					var message DeltaMessage
					if err := json.Unmarshal([]byte(data), &message); err != nil {
						fmt.Printf("Error unmarshalling DeltaMessage: %v\n", err)
						continue
					}
					// Log the value for each text entry in the content array
					for _, content := range message.Delta.Content {
						if content.Type == "text" {
							fmt.Println("Logged Value:", content.Text.Value)
							valueChannel <- content.Text.Value
						}
					}
				}
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading stream: %v\n", err)
		}
	}()

	return valueChannel, nil
}

func (g *gptAdaptorClient) GetRunDetails(threadID, runID string) (*RunResponse, error) {
	urls := fmt.Sprintf("%v/%v/runs/%v", ThreadEndpoint, threadID, runID)

	req, err := http.NewRequest("GET", urls, nil)
	if err != nil {
		return nil, err
	}

	// Sử dụng hàm helper để thêm headers chung
	g.addCommonHeaders(req)

	resp, err := g.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		var errMsg string
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			errMsg = "Failed to read response body"
		} else {
			errMsg = string(bodyBytes)
		}
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, errMsg)
	}
	// Xử lý response
	var response RunResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (g *gptAdaptorClient) UploadFile(filePath string) (*UploadFileResponse, error) {
	// Mở file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Tạo multipart form request
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	// Thêm field 'purpose'
	_ = writer.WriteField("purpose", "assistants")

	// Thêm file vào form
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("error creating form file: %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("Error copying file to form file: %v", err)
	}

	// Đóng writer
	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("Error closing writer: %v", err)
	}

	// Tạo và gửi request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/files", &buffer)
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %v", err)
	}
	req.Header.Add("Authorization", "Bearer "+g.ApiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Kiểm tra response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Upload failed with status code: %d", resp.StatusCode)
	}

	// Decode response
	var uploadResp UploadFileResponse
	err = json.NewDecoder(resp.Body).Decode(&uploadResp)
	if err != nil {
		return nil, fmt.Errorf("Error decoding response JSON: %v", err)
	}

	return &uploadResp, nil
}

func (g *gptAdaptorClient) WaitForRunCompletion(threadID, runID string) (*RunResponse, error) {
	timeout := time.NewTimer(2 * time.Minute) // Sets a timer for 2 minutes.
	ticker := time.NewTicker(5 * time.Second) // Checks every 5 seconds.
	defer func() {
		timeout.Stop()
		ticker.Stop()
	}()

	for {
		select {
		case <-timeout.C:
			fmt.Println("Timeout reached. Final status unknown.")
			return nil, fmt.Errorf("timeout reached. Final status unknown")

		case <-ticker.C:
			runResponse, err := g.GetRunDetails(threadID, runID)
			if err != nil {
				fmt.Println("Error fetching run details:", err)
				return nil, fmt.Errorf("error fetching run details: %v", err)
			}

			fmt.Printf("Checking run status: %s\n", runResponse.Status)

			if runResponse.Status == "completed" {
				fmt.Println("Run completed successfully.")
				return runResponse, nil
			} else if runResponse.Status == "failed" {
				fmt.Println("Run failed.")
				return runResponse, fmt.Errorf("run failed with status: %s", runResponse.Status)
			}
			// Continues to poll until the run completes or fails.
		}
	}
}
