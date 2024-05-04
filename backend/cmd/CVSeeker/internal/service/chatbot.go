package services

import (
	"CVSeeker/internal/ginLogger"
	"CVSeeker/internal/meta"
	"CVSeeker/pkg/gpt"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type IChatbotService interface {
	StartChatSession(c *gin.Context) (*meta.BasicResponse, error)
	SendMessageToChat(c *gin.Context, threadID, message string) (*meta.BasicResponse, error)
}

type ChatbotService struct {
	assistantClient gpt.IGptAdaptorClient
}

type ChatbotServiceArgs struct {
	dig.In
	AssistantClient gpt.IGptAdaptorClient
}

func NewChatbotService(args ChatbotServiceArgs) IChatbotService {
	return &ChatbotService{
		assistantClient: args.AssistantClient,
	}
}

func (_this *ChatbotService) StartChatSession(c *gin.Context) (*meta.BasicResponse, error) {
	assistantRequest := gpt.AssistantRequest{
		Model: "gpt-3.5-turbo",
	}
	assistant, err := _this.assistantClient.CreateAssistant(assistantRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to create gpt: %v", err)
	}

	thread, err := _this.assistantClient.CreateThread()
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to create thread: %v", err)
		return nil, err
	}

	return &meta.BasicResponse{
		Meta: meta.Meta{Code: 200},
		Data: map[string]string{
			"assistant_id": assistant.ID,
			"thread_id":    thread.ID,
		},
	}, nil
}

func (_this *ChatbotService) SendMessageToChat(c *gin.Context, threadID, message string) (*meta.BasicResponse, error) {
	messageRequest := gpt.CreateMessageRequest{
		Content: message,
		Role:    "user",
	}
	response, err := _this.assistantClient.CreateMessage(threadID, messageRequest)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to send message: %v", err)
		return nil, err
	}

	return &meta.BasicResponse{
		Meta: meta.Meta{Code: 200},
		Data: response,
	}, nil
}
