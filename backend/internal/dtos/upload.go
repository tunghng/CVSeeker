package dtos

type UploadDTO struct {
	DocumentID string `json:"documentId,omitempty"` // omitempty to not display if empty
	Status     string `json:"status"`
	Name       string `json:"name"`
	CreatedAt  int64  `json:"createdAt"` // Assuming date is formatted as a string for the client
}
