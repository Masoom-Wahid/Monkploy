package seeder

import (
	"platform/app/models"
	"platform/app/v1/services"
	"platform/database"
)

type AppSeeder interface {
	SeedApps()
}

type appSeeder struct {
	services services.Supplier
}

func NewAppSeeder(services services.Supplier) AppSeeder {
	return &appSeeder{
		services: services,
	}
}

func (u *appSeeder) SeedApps() {
	db := database.GetDB()

	user := models.App{
		Name: "Backend Api",
		Description: "Anthoer App",
	}

	db.Model(&user).Create(&user)
}
