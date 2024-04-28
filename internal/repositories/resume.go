package repositories

import (
	"CVSeeker/internal/models"
	"CVSeeker/pkg/db"
	"time"
)

type IResumeRepository interface {
	Create(db *db.DB, resume *models.Resume) (*models.Resume, error)
	Update(db *db.DB, resume *models.Resume) error
	FindByID(db *db.DB, resumeID int) (*models.Resume, error)
}

type resumeRepository struct{}

func NewResumeRepository() IResumeRepository {
	return &resumeRepository{}
}

func (_this *resumeRepository) Create(db *db.DB, resume *models.Resume) (*models.Resume, error) {
	if err := db.DB().Table(models.TableNameResume).Create(resume).Error; err != nil {
		return nil, err
	}
	return resume, nil
}

func (_this *resumeRepository) Update(db *db.DB, resume *models.Resume) error {
	resume.UpdatedAt = time.Now()
	return db.DB().Table(models.TableNameResume).Where("resume_id = ?", resume.ResumeID).Updates(resume).Error
}

func (_this *resumeRepository) FindByID(db *db.DB, resumeID int) (*models.Resume, error) {
	var resume models.Resume
	if err := db.DB().Table(models.TableNameResume).Where("resume_id = ?", resumeID).First(&resume).Error; err != nil {
		return nil, err
	}
	return &resume, nil
}
