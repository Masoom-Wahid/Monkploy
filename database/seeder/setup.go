package seeder

type SeederSupplier interface {
	Run()
	AppSeeder() AppSeeder
}

type seederSupplier struct {
}

func (seeder *seederSupplier) AppSeeder() AppSeeder {
	return NewAppSeeder()
}

func NewSeederSupplier() SeederSupplier {
	return &seederSupplier{}
}

func (seeder *seederSupplier) Run() {
	seeder.AppSeeder().SeedApps()
}
