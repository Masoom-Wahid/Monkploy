package middlewares

import (
	"platform/app/v1/services"
)

type middlewareSupplier struct {
	services services.Supplier
}

type MiddlewareSupplier interface {
}

func NewMiddlewareSupplier(services services.Supplier) MiddlewareSupplier {
	return &middlewareSupplier{
		services: services,
	}
}
