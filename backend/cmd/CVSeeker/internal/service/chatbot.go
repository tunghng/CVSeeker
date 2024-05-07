package services

import (
	"CVSeeker/cmd/CVSeeker/internal/cfg"
	"CVSeeker/internal/ginLogger"
	"CVSeeker/internal/meta"
	"CVSeeker/pkg/elasticsearch"
	"CVSeeker/pkg/gpt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"net/http"
	"strings"
)

type IChatbotService interface {
	StartChatSession(c *gin.Context, ids string) (*meta.BasicResponse, error)
	SendMessageToChat(c *gin.Context, threadID, message string) (*meta.BasicResponse, error)
	ListMessage(c *gin.Context, request gpt.ListMessageRequest) (*meta.BasicResponse, error)
}

type ChatbotService struct {
	assistantClient gpt.IGptAdaptorClient
	elasticClient   elasticsearch.IElasticsearchClient
}

type ChatbotServiceArgs struct {
	dig.In
	AssistantClient gpt.IGptAdaptorClient
	ElasticClient   elasticsearch.IElasticsearchClient
}

func NewChatbotService(args ChatbotServiceArgs) IChatbotService {
	return &ChatbotService{
		assistantClient: args.AssistantClient,
		elasticClient:   args.ElasticClient,
	}
}

func (_this *ChatbotService) StartChatSession(c *gin.Context, ids string) (*meta.BasicResponse, error) {
	// Parse the IDs from the string
	idArray := strings.Split(ids, ", ")

	// Fetch documents from Elasticsearch
	documents, err := _this.elasticClient.FetchDocumentsByIDs(c, "your_index_name", idArray)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to fetch documents: %v", err)
		return nil, err
	}

	// Format the documents' content
	var fullTextContent strings.Builder
	fullTextContent.WriteString("You will use these information to answer questions from the user: ")
	for _, doc := range documents {
		fullTextContent.WriteString(doc.Content + " ")
	}

	// Create the initial message for the thread
	initMessage := gpt.CreateMessageRequest{
		Role:    "user",
		Content: fullTextContent.String(),
	}

	// Create a new thread with the initial message
	threadRequest := gpt.CreateThreadRequest{
		Messages: []gpt.CreateMessageRequest{initMessage},
	}

	thread, err := _this.assistantClient.CreateThread(threadRequest)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to create thread: %v", err)
		return nil, err
	}

	// Prepare the response with the thread information
	response := &meta.BasicResponse{
		Meta: meta.Meta{
			Code:    200,
			Message: "Session started successfully with initial data",
		},
		Data: thread,
	}

	return response, nil
}

func (_this *ChatbotService) SendMessageToChat(c *gin.Context, threadID, message string) (*meta.BasicResponse, error) {
	DefaultAssistant := viper.GetString(cfg.DefaultAssistant)

	messageRequest := gpt.CreateMessageRequest{
		Content: message,
		Role:    "user",
	}

	_, err := _this.assistantClient.CreateMessage(threadID, messageRequest)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to send message: %v", err)
		return nil, err
	}

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
		return nil, err
	}

	//Get current list message
	listMessageResponse, err := _this.assistantClient.ListMessages(threadID, 2, "", "", "")
	if err != nil {
		ginLogger.Gin(c).Errorf("Error when list message: %v", err)
		return nil, err
	}
	if len(listMessageResponse.Data) == 0 {
		ginLogger.Gin(c).Errorf("Error when get list message: %v", err)
		return nil, err
	}

	response := &meta.BasicResponse{
		Meta: meta.Meta{
			Code:    200,
			Message: "Response retrieved successfully",
		},
		Data: listMessageResponse,
	}

	// Return the result of the completed run
	return response, nil
}

func (_this *ChatbotService) ListMessage(c *gin.Context, request gpt.ListMessageRequest) (*meta.BasicResponse, error) {
	resp, err := _this.assistantClient.ListMessages(request.ThreadId, request.Limit, request.Order, request.After, request.Before)
	if err != nil {
		ginLogger.Gin(c).Errorf("Error when create assistant: %v", err)
		return nil, err
	}
	response := &meta.BasicResponse{
		Meta: meta.Meta{
			Code:    http.StatusOK,
			Message: "Success",
		},
		Data: resp,
	}
	return response, nil
}
