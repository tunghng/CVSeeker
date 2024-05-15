package dtos

type QueryRequest struct {
	Content string `json:"content"`
}

type StartChatRequest struct {
	Ids        string `json:"ids"`
	ThreadName string `json:"threadName"`
}

type ResumesRequest struct {
	Resumes []ResumeData `json:"resumes"`
}

type ResumeData struct {
	Content   string `json:"content"`
	FileBytes string `json:"fileBytes"`
	Name      string `json:"name"`
}
