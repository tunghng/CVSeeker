package handlers

import (
	services "CVSeeker/cmd/CVSeeker/internal/service"
	"CVSeeker/internal/errors"
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
// @Description Starts a new chat session by creating an gpt and a thread.
// @Tags Chatbot
// @Accept json
// @Produce json
// @Success 200 {object} meta.BasicResponse
// @Failure 400,500 {object} meta.Error
// @Router /chat/start [POST]
func (_this *ChatbotHandler) StartChatSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := _this.chatbotService.StartChatSession(c)
		_this.HandleResponse(c, resp, err)
	}
}

// SendMessage
// @Summary Send a message to a chat session
// @Description Sends a message to the specified chat session.
// @Tags Chatbot
// @Accept json
// @Produce json
// @Param thread_id path string true "Thread ID"
// @Param content body string true "Message content"
// @Success 200 {object} meta.BasicResponse
// @Failure 400,500 {object} meta.Error
// @Router /chat/{thread_id}/send [POST]
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
