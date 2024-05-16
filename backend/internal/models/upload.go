package models

import (
	"time"
)

const TableNameUpload = "upload"

// Upload represents the schema of the "upload_history" table.
type Upload struct {
	ID         int       `gorm:"column:id;primary_key;auto_increment" json:"id"`
	DocumentID string    `gorm:"column:document_id;type:varchar(255)" json:"documentId"`
	Status     string    `gorm:"column:status;type:varchar(100)" json:"status"`
	Name       string    `gorm:"column:name;type:varchar(255)" json:"name"`
	UUID       string    `gorm:"column:uuid;type:varchar(255)" json:"uuid"`
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP" json:"updatedAt"`
}

// TableName overrides the table name used by Upload to `upload_history`
func (Upload) TableName() string {
	return TableNameUpload
}
