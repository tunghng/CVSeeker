package models

import (
	"time"
)

const TableNameThread = "threads"

// Thread represents a conversation or interaction session.
type Thread struct {
	ID        string    `gorm:"primaryKey;type:varchar(100)" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"` // Added name field
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (Thread) TableName() string {
	return TableNameThread
}
