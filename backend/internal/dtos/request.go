package dtos

type QueryRequest struct {
	Content string `json:"content"`
}

type StartChatRequest struct {
	Ids        string `json:"ids"`
	ThreadName string `json:"threadName"`
}

type ResumeRequest struct {
	Content   string `json:"content"`
	FileBytes string `json:"fileBytes"` // base64 encoded string of the file
}
