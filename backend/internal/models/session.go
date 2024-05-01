package models

import (
	"time"
)

const TableNameSessions = "sessions"

type Session struct {
	SessionID int       `gorm:"column:session_id;PRIMARY_KEY;AUTO_INCREMENT" json:"sessionId"`
	StartTime time.Time `gorm:"column:start_time" json:"startTime"`
	EndTime   time.Time `gorm:"column:end_time" json:"endTime"`
	Resumes   []Resume  `gorm:"many2many:session_resumes;" json:"resumes"`
}

func (Session) TableName() string {
	return TableNameSessions
}
