package migrations

import (
	"platform/app/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	db.AutoMigrate(&models.App{})
}
