package repositories

import (
	"platform/app/models"
	"platform/database"

	"gorm.io/gorm"
)

type AppRepository interface {
	List() ([]models.App, error)
	Create(app *models.App) *gorm.DB
}

type appRepository struct {
}

func NewAppRepository() *appRepository {
	return &appRepository{}
}

func (r *appRepository) List() ([]models.App, error) {
	db := database.GetDB()
	var apps []models.App
	err := db.Model(&models.App{}).Order("id asc").Find(&apps).Error
	return apps, err
}

func (r *appRepository) Create(app *models.App) *gorm.DB {
	db := database.GetDB()

	return db.Model(&models.App{}).Create(app)

}
