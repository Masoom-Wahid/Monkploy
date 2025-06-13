package seeder

import (
	"platform/app/v1/services"
)

type SeederSupplier interface {
	Run()
	AppSeeder() AppSeeder
}

type seederSupplier struct {
	services services.Supplier
}

func (seeder *seederSupplier) AppSeeder() AppSeeder {
	return NewAppSeeder(seeder.services)
}

func NewSeederSupplier(services services.Supplier) SeederSupplier {
	return &seederSupplier{
		services: services,
	}
}

func (seeder *seederSupplier) Run() {
	seeder.AppSeeder().SeedApps()
}


