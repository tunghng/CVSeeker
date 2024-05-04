package gpt

const (
	AssistantsSuffix      = "/assistants"
	AssistantsFilesSuffix = "/files"
	OpenaiAssistantsV1    = "assistants=v1"

	AssistantEndpoint = "https://api.openai.com/v1/assistants"
	ThreadEndpoint    = "https://api.openai.com/v1/threads"
)

type Assistant struct {
	ID           string          `json:"id"`
	Object       string          `json:"object"`
	CreatedAt    int64           `json:"created_at"`
	Name         *string         `json:"name,omitempty"`
	Description  *string         `json:"description,omitempty"`
	Model        string          `json:"model"`
	Instructions *string         `json:"instructions,omitempty"`
	Tools        []AssistantTool `json:"tools,omitempty"`
	FileIDs      []string        `json:"file_ids,omitempty"`
	Metadata     map[string]any  `json:"metadata,omitempty"`
}

type AssistantToolType string

const (
	AssistantToolTypeCodeInterpreter AssistantToolType = "code_interpreter"
	AssistantToolTypeRetrieval       AssistantToolType = "retrieval"
	AssistantToolTypeFunction        AssistantToolType = "function"
)

type AssistantTool struct {
	Id       string              `json:"id"`
	Type     AssistantToolType   `json:"type"`
	Function *FunctionDefinition `json:"function,omitempty"`
}

type AssistantRequest struct {
	Model        string          `json:"model"`
	Name         *string         `json:"name,omitempty"`
	Description  *string         `json:"description,omitempty"`
	Instructions *string         `json:"instructions,omitempty"`
	Tools        []AssistantTool `json:"tools,omitempty"`
	FileIDs      []string        `json:"file_ids,omitempty"`
	Metadata     map[string]any  `json:"metadata,omitempty"`
}

type AssistantResponse struct {
	ID           string                 `json:"id"`
	Object       string                 `json:"object"`
	CreatedAt    int64                  `json:"created_at"`
	Name         string                 `json:"name"`
	Description  *string                `json:"description"`
	Model        string                 `json:"model"`
	Instructions string                 `json:"instructions"`
	Tools        []AssistantTool        `json:"tools"`
	FileIDs      []string               `json:"file_ids"`
	Metadata     map[string]interface{} `json:"metadata"`
}

// AssistantsList is a list of assistants.
type AssistantsList struct {
	Assistants []Assistant `json:"data"`
	LastID     *string     `json:"last_id"`
	FirstID    *string     `json:"first_id"`
	HasMore    bool        `json:"has_more"`
}

type AssistantDeleteResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

type AssistantFile struct {
	ID          string `json:"id"`
	Object      string `json:"object"`
	CreatedAt   int64  `json:"created_at"`
	AssistantID string `json:"assistant_id"`
}

type AssistantFileRequest struct {
	FileID string `json:"file_id"`
}

type FunctionDefinition struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Parameters  any    `json:"parameters"`
}

type ThreadResponse struct {
	ID        string                 `json:"id"`
	Object    string                 `json:"object"`
	CreatedAt int64                  `json:"created_at"`
	Metadata  map[string]interface{} `json:"metadata"`
}

type DeleteThreadResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

type CreateMessageRequest struct {
	Role    string   `json:"role"`
	Content string   `json:"content"`
	FileIds []string `json:"file_ids,omitempty"`
}

type MessageContent struct {
	Type string `json:"type"`
	Text struct {
		Value       string        `json:"value"`
		Annotations []interface{} `json:"annotations"`
	} `json:"text"`
}

type MessageResponse struct {
	ID          string                 `json:"id"`
	Object      string                 `json:"object"`
	CreatedAt   int64                  `json:"created_at"`
	ThreadID    string                 `json:"thread_id"`
	Role        string                 `json:"role"`
	Content     []MessageContent       `json:"content"`
	FileIDs     []string               `json:"file_ids"`
	AssistantID *string                `json:"assistant_id"`
	RunID       *string                `json:"run_id"`
	Metadata    map[string]interface{} `json:"metadata"`
}

type ListMessagesResponse struct {
	Object  string            `json:"object"`
	Data    []MessageResponse `json:"data"`
	FirstID string            `json:"first_id"`
	LastID  string            `json:"last_id"`
	HasMore bool              `json:"has_more"`
}
type CheckRunRequest struct {
	ThreadId string `json:"threadId"`
	RuneId   string `json:"runeId"`
}

type ListMessageRequest struct {
	ThreadId string `json:"threadId"`
	Limit    int    `json:"limit"`
	Order    string `json:"order"`
	After    string `json:"after"`
	Before   string `json:"before"`
}

type CreateRunRequest struct {
	AssistantID  string            `json:"assistant_id"`
	Model        *string           `json:"model,omitempty"`
	Instructions *string           `json:"instructions,omitempty"`
	Tools        []AssistantTool   `json:"tools,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}

type RunResponse struct {
	ID             string                 `json:"id"`
	Object         string                 `json:"object"`
	CreatedAt      int64                  `json:"created_at"`
	AssistantID    string                 `json:"assistant_id"`
	ThreadID       string                 `json:"thread_id"`
	Status         string                 `json:"status"`
	StartedAt      *int64                 `json:"started_at,omitempty"`
	ExpiresAt      *int64                 `json:"expires_at,omitempty"`
	CancelledAt    *int64                 `json:"cancelled_at,omitempty"`
	FailedAt       *int64                 `json:"failed_at,omitempty"`
	CompletedAt    *int64                 `json:"completed_at,omitempty"`
	RequiredAction *RequiredAction        `json:"required_action,omitempty"`
	LastError      *string                `json:"last_error,omitempty"`
	Model          string                 `json:"model"`
	Instructions   *string                `json:"instructions,omitempty"`
	Tools          []AssistantTool        `json:"tools"`
	FileIDs        []string               `json:"file_ids"`
	Metadata       map[string]interface{} `json:"metadata"`
}

type ToolOutput struct {
	ToolCallID string `json:"tool_call_id"`
	Output     string `json:"output"`
}

type SubmitToolOutputsRequest struct {
	ToolOutputs []ToolOutput `json:"tool_outputs"`
}

type FunctionArguments struct {
	QuestionID int `json:"questionID"`
}

type ToolCallFunction struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type ToolCall struct {
	ID       string           `json:"id"`
	Type     string           `json:"type"`
	Function ToolCallFunction `json:"function"`
}

type RequiredAction struct {
	Type              string `json:"type"`
	SubmitToolOutputs struct {
		ToolCalls []ToolCall `json:"tool_calls"`
	} `json:"submit_tool_outputs"`
}

type UploadFileResponse struct {
	ID          string `json:"id"`
	Object      string `json:"object"`
	CreatedAt   int64  `json:"created_at"`
	AssistantID string `json:"assistant_id"`
}
