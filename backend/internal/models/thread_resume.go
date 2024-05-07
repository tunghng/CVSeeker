package models

import (
	"time"
)

const TableNameThreadResume = "thread_resume"

type ThreadResume struct {
	ThreadId  string    `gorm:"column:thread_id;primaryKey" json:"threadId"`
	ResumeId  int       `gorm:"column:resume_id;primaryKey" json:"resumeId"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (ThreadResume) TableName() string {
	return TableNameThreadResume
}
