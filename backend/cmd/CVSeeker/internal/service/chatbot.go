package services

import (
	"CVSeeker/internal/ginLogger"
	"CVSeeker/internal/meta"
	"CVSeeker/pkg/gpt"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

const (
	DefaultAssistant = "asst_zIkxsuNW2nhjWJZhUiTV6Vp9"
	RoleUser         = "user"
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
	thread, err := _this.assistantClient.CreateThread()
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to create thread: %v", err)
		return nil, err
	}

	return &meta.BasicResponse{
		Meta: meta.Meta{Code: 200},
		Data: map[string]string{
			"thread_id": thread.ID,
		},
	}, nil
}

func (_this *ChatbotService) SendMessageToChat(c *gin.Context, threadID, message string) (*meta.BasicResponse, error) {
	// Create a message with the user's input
	messageRequest := gpt.CreateMessageRequest{
		Content: message,
		Role:    RoleUser,
	}

	_, err := _this.assistantClient.CreateMessage(threadID, messageRequest)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to send message: %v", err)
		return nil, err
	}

	// Create a run to process the user's message
	runRequest := gpt.CreateRunRequest{
		AssistantID: DefaultAssistant,
	}
	runResponse, err := _this.assistantClient.CreateRun(threadID, runRequest)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to create run: %v", err)
		return nil, err
	}

	// Wait for the completion of the run to get the response from the assistant
	completedRun, err := _this.assistantClient.WaitForRunCompletion(threadID, runResponse.ID)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to wait for run completion: %v", err)
		return nil, err
	}

	if completedRun.Status != "completed" {
		ginLogger.Gin(c).Errorf("run did not complete successfully")
		return nil, fmt.Errorf("run did not complete successfully")
	}

	// Return the result of the completed run
	return &meta.BasicResponse{
		Meta: meta.Meta{Code: 200},
		Data: completedRun,
	}, nil
}
