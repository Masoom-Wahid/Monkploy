package repositories

import (
	"platform/app/models"
	"platform/database"

	"gorm.io/gorm"
)

type ServiceRepository interface {
	List(appId string) ([]models.Service, error)
	Create(s *models.Service, appId string) (*gorm.DB, error)
	Delete(id string, appId string) error
}

type serviceRepository struct {
}

func NewServiceRepository() *serviceRepository {
	return &serviceRepository{}
}

func (s *serviceRepository) List(appId string) ([]models.Service, error) {
	db := database.GetDB()
	var services []models.Service
	err := db.Model(&models.Service{}).Where("app_id = ?", appId).Order("id asc").Find(&services).Error
	return services, err
}

func (r *serviceRepository) Create(s *models.Service, appId string) (*gorm.DB, error) {
	db := database.GetDB()

	// Check if the app exists
	var app models.App
	if err := db.Where("id = ?", appId).First(&app).Error; err != nil {
		return nil, err
	}

	s.AppID = app.ID

	result := db.Model(&models.Service{}).Create(s)
	
	return result, result.Error
}

func (r *serviceRepository) Delete(id string, appId string) error {
	db := database.GetDB()
	

	var app models.App
	if err := db.Where("id = ?", appId).First(&app).Error; err != nil {
		return err
	}

	return db.Where("id = ?", id).Where("app_id = ?", appId).Delete(&models.Service{}).Error
}
