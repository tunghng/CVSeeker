package models

import (
	"time"
)

const TableNameThreadResume = "thread_resumes"

// ThreadResume represents the relationship between a thread and resumes.
type ThreadResume struct {
	ThreadID  string    `gorm:"primaryKey;column:thread_id;type:varchar(100)" json:"threadId"`
	ResumeID  string    `gorm:"primaryKey;column:resume_id" json:"resumeId"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (ThreadResume) TableName() string {
	return TableNameThreadResume
}
