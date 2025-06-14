package requests

import "platform/app/models"

type ServiceRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Type string `json:"type"`
}

type ServiceResponse struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description"`
	Type string `json:"type"`

}

func (r *ServiceRequest) ToServiceModel() models.Service {
	service := models.Service{
		Name:        r.Name,
		Description: r.Description,
		Type: r.Type,
	}

	return service
}
