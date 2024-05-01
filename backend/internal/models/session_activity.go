package models

import (
	"time"
)

const TableNameSessionActivities = "session_activities"

type SessionActivity struct {
	ActivityID   int       `gorm:"column:activity_id;PRIMARY_KEY;AUTO_INCREMENT" json:"activityId"`
	SessionID    int       `gorm:"column:session_id" json:"sessionId"`
	QueryText    string    `gorm:"column:query_text;type:text" json:"queryText"`
	ResponseText string    `gorm:"column:response_text;type:text" json:"responseText"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (SessionActivity) TableName() string {
	return TableNameSessionActivities
}
