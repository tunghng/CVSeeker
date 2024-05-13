package repositories

import (
	"CVSeeker/internal/models"
	"CVSeeker/pkg/db"
)

// IUploadRepository defines the interface for the upload repository.
type IUploadRepository interface {
	Create(db *db.DB, upload *models.Upload) (*models.Upload, error)
	GetAll(db *db.DB) ([]models.Upload, error)
	Update(db *db.DB, upload *models.Upload) error
}

// uploadRepository implements the IUploadRepository interface.
type uploadRepository struct{}

// NewUploadRepository creates a new instance of uploadRepository.
func NewUploadRepository() IUploadRepository {
	return &uploadRepository{}
}

// Create inserts a new upload record into the database.
func (_this *uploadRepository) Create(db *db.DB, upload *models.Upload) (*models.Upload, error) {
	if err := db.DB().Table(models.TableNameUpload).Create(upload).Error; err != nil {
		return nil, err
	}
	return upload, nil
}

// GetAll retrieves all upload records from the database, sorted from latest to oldest.
func (_this *uploadRepository) GetAll(db *db.DB) ([]models.Upload, error) {
	var uploads []models.Upload
	if err := db.DB().Table(models.TableNameUpload).Order("created_at DESC").Find(&uploads).Error; err != nil {
		return nil, err
	}
	return uploads, nil
}

func (_this *uploadRepository) Update(db *db.DB, upload *models.Upload) error {
	return db.DB().Table(models.TableNameUpload).Where("id = ?", upload.ID).Updates(upload).Error
}
