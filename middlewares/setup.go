package middlewares

type middlewareSupplier struct {
}

type MiddlewareSupplier interface {
}

func NewMiddlewareSupplier() MiddlewareSupplier {
	return &middlewareSupplier{}
}
