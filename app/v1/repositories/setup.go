package repositories

import "platform/config"

type RepoSupplier interface {
	AppRepository() AppRepository
	ServiceRepository() ServiceRepository
}

type repoSupplier struct {
	appConfig config.Config
}

func NewRepoSupplier(config config.Config) RepoSupplier {
	return &repoSupplier{
		appConfig: config,
	}
}

func (r *repoSupplier) AppRepository() AppRepository {
	return NewAppRepository()
}


func (r *repoSupplier) ServiceRepository() ServiceRepository {
	return NewServiceRepository()
}