package models

import (
	"time"
)

const TableNameResume = "resumes"

type Resume struct {
	ResumeID        int       `gorm:"column:resume_id;PRIMARY_KEY;AUTO_INCREMENT" json:"resumeId"`
	FullText        string    `gorm:"column:full_text;type:text" json:"fullText"`
	DownloadLink    string    `gorm:"column:download_link" json:"downloadLink"`
	VectorEmbedding string    `gorm:"column:vector_embedding;type:text" json:"vectorEmbedding"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (Resume) TableName() string {
	return TableNameResume
}
