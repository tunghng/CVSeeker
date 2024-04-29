package elasticsearch

import "time"

type ResumeDTO struct {
	ResumeID        string    `json:"resume_id"`
	FullText        string    `json:"full_text"`
	DownloadLink    string    `json:"download_link"`
	VectorEmbedding string    `json:"vector_embedding"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type ElasticResponse struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	URL     string `json:"url"`
}
