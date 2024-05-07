package repositories

import (
	"CVSeeker/internal/models"
	"CVSeeker/pkg/db"
	"time"
)

type IThreadRepository interface {
	Create(db *db.DB, thread *models.Thread) (*models.Thread, error)
	Update(db *db.DB, thread *models.Thread) error
	FindByID(db *db.DB, threadID string) (*models.Thread, error)
	GetAllThreadIDs(db *db.DB) ([]string, error)
}

type threadRepository struct{}

func NewThreadRepository() IThreadRepository {
	return &threadRepository{}
}

func (_this *threadRepository) Create(db *db.DB, thread *models.Thread) (*models.Thread, error) {
	thread.CreatedAt = time.Now() // Set creation time
	thread.UpdatedAt = time.Now() // Set update time
	if err := db.DB().Table(models.TableNameThread).Create(thread).Error; err != nil {
		return nil, err
	}
	return thread, nil
}

func (_this *threadRepository) Update(db *db.DB, thread *models.Thread) error {
	thread.UpdatedAt = time.Now() // Update the modified time
	return db.DB().Table(models.TableNameThread).Where("id = ?", thread.ID).Updates(thread).Error
}

func (_this *threadRepository) FindByID(db *db.DB, threadID string) (*models.Thread, error) {
	var thread models.Thread
	if err := db.DB().Table(models.TableNameThread).Where("id = ?", threadID).First(&thread).Error; err != nil {
		return nil, err
	}
	return &thread, nil
}

func (_this *threadRepository) GetAllThreadIDs(db *db.DB) ([]string, error) {
	var ids []string
	if err := db.DB().Table(models.TableNameThread).Pluck("id", &ids).Error; err != nil {
		return nil, err
	}
	return ids, nil
}
