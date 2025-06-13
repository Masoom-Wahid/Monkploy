package repositories

import "platform/config"

type DAO interface {

}

type dao struct {
	// db *gorm.DB
	appConfig config.Config
}

func NewDAO(config config.Config) DAO {
	return &dao{
		appConfig: config,
	}
}
