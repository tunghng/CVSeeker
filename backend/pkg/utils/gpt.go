package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
	apiKey      = "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" // Thay thế bằng API Key của bạn
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatPayload struct {
	Messages []Message `json:"messages"`
	Model    string    `json:"model"`
}

type ChatResponse struct {
	Id      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

func AskChatGptForSummarizeVideo(title, transcriptJson, t string) (string, error) {
	prompt := generatePromptVideoSummarize(title, transcriptJson, t)

	messages := []Message{
		{
			Role:    "user",
			Content: prompt,
		},
	}

	payload := ChatPayload{
		Messages: messages,
		Model:    "summarizer-3.5-turbo",
	}

	response, err := callChatGPTAPI(payload)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	}
	return "", nil
}

func AskChatGptToVerifyTranslation(original, translation, t string) (string, error) {
	prompt := generatePrompt(original, translation, t)

	messages := []Message{
		{
			Role:    "user",
			Content: prompt,
		},
	}

	payload := ChatPayload{
		Messages: messages,
		Model:    "summarizer-3.5-turbo",
	}

	response, err := callChatGPTAPI(payload)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	}
	return "", nil
}

func generatePromptVideoSummarize(title, transcriptJson, types string) string {
	return fmt.Sprintf(`Summarize the following into 5-10 bullet points in the format "start time": "point"
Title: "%v"
Transcript json type: %v`, title, transcriptJson)

}

func generatePrompt(original, translation, types string) string {
	if types == "E2V" {
		return fmt.Sprintf(`Bạn là một giảng viên chuyên ngành tiếng anh Y Khoa, hãy giúp tôi kiểm tra bản dịch dưới đây của tôi dịch từ Tiếng Anh sang Tiếng Việt và nhận xét ngắn gọn trong 3 câu.
Bài viết tiếng anh: %s .
Bản dịch của tôi: %s`, original, translation)
	} else {
		return fmt.Sprintf(`Bạn là một giảng viên chuyên ngành tiếng anh Y Khoa, hãy giúp tôi kiểm tra bản dịch dưới đây của tôi dịch từ Tiếng Việt sang Tiếng Anh và nhận xét ngắn gọn trong 3 câu.
Bài viết Tiếng Việt: %s .
Bản dịch Tiếng Anh của tôi: %s`, original, translation)
	}
}

func callChatGPTAPI(payload ChatPayload) (*ChatResponse, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var chatResponse ChatResponse
	err = json.Unmarshal(body, &chatResponse)
	if err != nil {
		return nil, err
	}

	return &chatResponse, nil
}
