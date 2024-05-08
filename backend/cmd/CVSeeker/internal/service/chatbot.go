package services

import (
	"CVSeeker/cmd/CVSeeker/internal/cfg"
	"CVSeeker/internal/dtos"
	"CVSeeker/internal/ginLogger"
	"CVSeeker/internal/meta"
	"CVSeeker/internal/models"
	"CVSeeker/internal/repositories"
	"CVSeeker/pkg/db"
	"CVSeeker/pkg/elasticsearch"
	"CVSeeker/pkg/gpt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"net/http"
	"strings"
)

type IChatbotService interface {
	StartChatSession(c *gin.Context, ids string, threadName string) (*meta.BasicResponse, error)
	SendMessageToChat(c *gin.Context, threadID, message string) (*meta.BasicResponse, error)
	ListMessage(c *gin.Context, request gpt.ListMessageRequest) (*meta.BasicResponse, error)
	GetAllThreads(c *gin.Context) (*meta.BasicResponse, error)
	GetResumesByThreadID(c *gin.Context, threadID string) (*meta.BasicResponse, error)
}

type ChatbotService struct {
	db               *db.DB
	assistantClient  gpt.IGptAdaptorClient
	elasticClient    elasticsearch.IElasticsearchClient
	threadRepo       repositories.IThreadRepository
	threadResumeRepo repositories.IThreadResumeRepository
}

type ChatbotServiceArgs struct {
	dig.In
	DB               *db.DB `name:"talentAcquisitionDB"`
	AssistantClient  gpt.IGptAdaptorClient
	ElasticClient    elasticsearch.IElasticsearchClient
	ThreadRepo       repositories.IThreadRepository
	ThreadResumeRepo repositories.IThreadResumeRepository
}

func NewChatbotService(args ChatbotServiceArgs) IChatbotService {
	return &ChatbotService{
		db:               args.DB,
		assistantClient:  args.AssistantClient,
		elasticClient:    args.ElasticClient,
		threadRepo:       args.ThreadRepo,
		threadResumeRepo: args.ThreadResumeRepo,
	}
}

func (_this *ChatbotService) StartChatSession(c *gin.Context, ids string, threadName string) (*meta.BasicResponse, error) {
	elasticDocumentName := viper.GetString(cfg.ElasticsearchDocumentIndex)

	// Parse the IDs from the string
	idArray := strings.Split(ids, ", ")

	// Fetch documents from Elasticsearch
	documents, err := _this.elasticClient.FetchDocumentsByIDs(c, elasticDocumentName, idArray)
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

	// Create a new thread instance in the database
	newThread := &models.Thread{
		ID:   thread.ID,
		Name: threadName,
	}

	_, err = _this.threadRepo.Create(_this.db, newThread)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to create new thread record: %v", err)
		return nil, err
	}

	var threadResumes []models.ThreadResume

	// Create new thread_resume instances in the database
	for _, id := range idArray {
		threadResume := models.ThreadResume{
			ThreadID: thread.ID,
			ResumeID: id,
		}
		threadResumes = append(threadResumes, threadResume)
	}
	err = _this.threadResumeRepo.CreateBulkThreadResume(_this.db, threadResumes)

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
	DefaultAssistant := viper.GetString(cfg.DefaultOpenAIAssistant)

	// Create message and add to thread
	messageRequest := gpt.CreateMessageRequest{
		Content: message,
		Role:    "user",
	}

	_, err := _this.assistantClient.CreateMessage(threadID, messageRequest)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to send message: %v", err)
		return nil, err
	}

	// Create run for assistant and thread
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

	// Update the 'UpdatedAt' column for the thread
	if err := _this.threadRepo.UpdateUpdatedAt(_this.db, threadID); err != nil {
		ginLogger.Gin(c).Errorf("failed to update thread: %v", err)
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

func (_this *ChatbotService) GetAllThreads(c *gin.Context) (*meta.BasicResponse, error) {
	modelThreads, err := _this.threadRepo.GetAllThreads(_this.db)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to get all threads: %v", err)
		return nil, err
	}

	// Map model threads to DTOs
	threadDTOs := make([]dtos.Thread, len(modelThreads))
	for i, modelThread := range modelThreads {
		threadDTOs[i] = dtos.Thread{
			ID:        modelThread.ID,
			Name:      modelThread.Name,
			UpdatedAt: modelThread.UpdatedAt.Unix(),
		}
	}

	response := &meta.BasicResponse{
		Meta: meta.Meta{
			Code:    200,
			Message: "All threads retrieved successfully",
		},
		Data: threadDTOs,
	}
	return response, nil
}

func (_this *ChatbotService) GetResumesByThreadID(c *gin.Context, threadID string) (*meta.BasicResponse, error) {
	elasticDocumentName := viper.GetString(cfg.ElasticsearchDocumentIndex)

	resumeIDs, err := _this.threadResumeRepo.GetResumeIDsByThreadID(_this.db, threadID)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to fetch resume IDs by thread ID: %v", err)
		return nil, err
	}

	// Fetch documents from Elasticsearch
	documents, err := _this.elasticClient.FetchDocumentsByIDs(c, elasticDocumentName, resumeIDs)
	if err != nil {
		ginLogger.Gin(c).Errorf("failed to fetch documents: %v", err)
		return nil, err
	}

	response := &meta.BasicResponse{
		Meta: meta.Meta{
			Code:    200,
			Message: "Resume IDs retrieved successfully for the thread",
		},
		Data: documents,
	}
	return response, nil
}
