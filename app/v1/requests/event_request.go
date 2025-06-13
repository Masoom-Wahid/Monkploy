package requests

type EventRequest struct {
	Type       string `json:"type" validate:"required"`
	Uuid       string `json:"uuid" validate:"required"`
	DbName     string `json:"db_name" validate:"required"`
	DbPassword string `json:"db_password" validate:"required"`
	Domain     string `json:"domain" validate:"required"`
}

type EventResponse struct {
	Type       string `json:"type" validate:"required"`
	Uuid       string `json:"uuid" validate:"required"`
	Status     string `json:"status" validate:"required"`
	Domain     string `json:"domain" validate:"required"`
	FullDomain string `json:"full_domain" validate:"required"`
	Detail     string `json:"detail" validate:"required"`
}
