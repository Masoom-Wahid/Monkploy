package seeder

import (
	"platform/app/models"
	"platform/database"
)

type AppSeeder interface {
	SeedApps()
}

type appSeeder struct {
}

func NewAppSeeder() AppSeeder {
	return &appSeeder{}
}

func (u *appSeeder) SeedApps() {
	db := database.GetDB()

	testApp := models.App{
		Name:        "Backend Api",
		Description: "Anthoer App",
	}

	db.Model(&models.App{}).Create(&testApp)
}
