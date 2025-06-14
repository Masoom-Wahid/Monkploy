package requests

import "platform/app/models"

type AppRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type AppResponse struct {
	Name string `json:"name" validate:"required"`

	Description string `json:"description"`
}

func (r *AppRequest) ToAppModel() models.App {
	app := models.App{
		Name:        r.Name,
		Description: r.Description,
	}

	return app
}
