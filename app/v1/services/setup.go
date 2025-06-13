package services

import (
	"platform/app/v1/repositories"
	"platform/config"
)

type Supplier interface {

}



func NewServiceSupplier(dao repositories.DAO,
	config config.Config,
) Supplier {
	return nil
}
