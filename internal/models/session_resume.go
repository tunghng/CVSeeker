package models

const TableNameSessionResumes = "session_resumes"

type SessionResume struct {
	SessionID int `gorm:"column:session_id;primary_key" json:"sessionId"`
	ResumeID  int `gorm:"column:resume_id;primary_key" json:"resumeId"`
}

func (SessionResume) TableName() string {
	return TableNameSessionResumes
}
