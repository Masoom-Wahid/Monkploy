package controllers

import (
	"platform/app/v1/services"
)

type ControllerSupplier interface {

}

type controllerSupplier struct {
	services services.Supplier
}

func NewControllerSupplier(services services.Supplier) ControllerSupplier {
	return &controllerSupplier{
		services: services,
	}
}

