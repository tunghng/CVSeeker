package repositories

import (
	"CVSeeker/internal/models"
	"CVSeeker/pkg/db"
	"github.com/jinzhu/gorm"
)

type IThreadResumeRepository interface {
	Create(db *db.DB, threadResume *models.ThreadResume) error
	CreateBulkThreadResume(db *db.DB, threadResumes []models.ThreadResume) error
	GetResumeIDsByThreadID(db *db.DB, threadID string) ([]string, error)
}

type threadResumeRepository struct{}

func NewThreadResumeRepository() IThreadResumeRepository {
	return &threadResumeRepository{}
}

func (_this *threadResumeRepository) Create(db *db.DB, threadResume *models.ThreadResume) error {
	return db.DB().Table(models.TableNameThreadResume).Create(threadResume).Error
}

func (_this *threadResumeRepository) CreateBulkThreadResume(db *db.DB, threadResumes []models.ThreadResume) error {
	return db.DB().Transaction(func(tx *gorm.DB) error {
		for _, threadResume := range threadResumes {
			if err := tx.Create(&threadResume).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (_this *threadResumeRepository) GetResumeIDsByThreadID(db *db.DB, threadID string) ([]string, error) {
	var ids []string
	if err := db.DB().Table(models.TableNameThreadResume).Where("thread_id = ?", threadID).Pluck("resume_id", &ids).Error; err != nil {
		return nil, err
	}
	return ids, nil
}
