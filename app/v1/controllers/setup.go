package controllers

import "platform/app/v1/repositories"

type ControllerSupplier interface {
	AppController() AppController
}

type controllerSupplier struct {
	repoSupplier *repositories.RepoSupplier
}

func NewControllerSupplier(repoSupplier *repositories.RepoSupplier) ControllerSupplier {
	return &controllerSupplier{repoSupplier: repoSupplier}
}

func (c *controllerSupplier) AppController() AppController {
	return NewAppController(*c.repoSupplier)
}
