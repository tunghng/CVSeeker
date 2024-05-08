package dtos

type QueryRequest struct {
	Content string `json:"content"`
}

type IdsRequest struct {
	Ids string `json:"ids"`
}

type ResumeRequest struct {
	Content   string `json:"content"`
	FileBytes string `json:"filebytes"` // base64 encoded string of the file
}
