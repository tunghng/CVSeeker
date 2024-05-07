package handlers

import (
	services "CVSeeker/cmd/CVSeeker/internal/service"
	"CVSeeker/cmd/CVSeeker/pkg/utils"
	"CVSeeker/internal/errors"
	"CVSeeker/pkg/gpt"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"strings"
)

type ChatbotHandler struct {
	BaseHandler
	chatbotService services.IChatbotService
}

type ChatbotHandlerParams struct {
	dig.In
	BaseHandler    BaseHandler
	ChatbotService services.IChatbotService
}

func NewChatbotHandler(params ChatbotHandlerParams) *ChatbotHandler {
	return &ChatbotHandler{
		BaseHandler:    params.BaseHandler,
		chatbotService: params.ChatbotService,
	}
}

// StartChatSession
// @Summary Start a new chat session
// @Description Starts a new chat session by creating an assistant and a thread, using specified documents.
// @Tags Chatbot
// @Accept json
// @Produce json
// @Param ids query string true "Comma-separated list of document IDs"
// @Success 200 {object} meta.BasicResponse
// @Failure 400,500 {object} meta.Error
// @Router /cvseeker/resumes/thread/start [POST]
func (_this *ChatbotHandler) StartChatSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		ids := c.Query("ids")
		if ids == "" {
			_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInvalidRequest))
			return
		}

		resp, err := _this.chatbotService.StartChatSession(c, ids)
		_this.HandleResponse(c, resp, err)
	}
}

// SendMessage
// @Summary Send a message to a chat session
// @Description Sends a message to the specified chat session.
// @Tags Chatbot
// @Accept json
// @Produce json
// @Param threadId path string true "Thread ID"
// @Param content query string true "Message content"
// @Success 200 {object} meta.BasicResponse
// @Failure 400,500 {object} meta.Error
// @Router /cvseeker/resumes/thread/{threadId}/send [POST]
func (_this *ChatbotHandler) SendMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		threadID := strings.TrimSpace(c.Param("threadId"))
		if threadID == "" {
			_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInvalidRequest))
			return
		}

		message := strings.TrimSpace(c.Query("content"))
		if message == "" {
			_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInvalidRequest))
			return
		}

		resp, err := _this.chatbotService.SendMessageToChat(c, threadID, message)
		_this.HandleResponse(c, resp, err)
	}
}

// ListMessage
// @Summary List messages belonging to a thread
// @Description Get a list of messages for a thread.
// @Tags Chatbot
// @Accept json
// @Produce json
// @Param threadId path string true "Thread ID"
// @Success  200  {object}  meta.BasicResponse
// @Failure   400,401,404,500  {object}  meta.Error
// @Security  BearerAuth
// @Router /cvseeker/resumes/thread/{threadId}/messages [GET]
func (_this *ChatbotHandler) ListMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		threadId := strings.TrimSpace(c.Param("threadId"))
		if threadId == "" {
			_this.RespondError(c, errors.NewCusErr(errors.ErrCommonInvalidRequest))
			return
		}
		limit := utils.Str2StrInt64(c.Query("limit"), true)
		after := strings.TrimSpace(c.Query("after"))
		before := strings.TrimSpace(c.Query("before"))

		var request gpt.ListMessageRequest
		request.ThreadId = threadId
		request.Limit = int(limit)
		request.After = after
		request.Before = before
		resp, err := _this.chatbotService.ListMessage(c, request)
		_this.HandleResponse(c, resp, err)
	}
}
