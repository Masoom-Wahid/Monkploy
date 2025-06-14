package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type App struct {
	ID        uuid.UUID      `gorm:"primaryKey" json:"id,omitempty"`
	Name        string         `gorm:"not null" json:"name,omitempty"`
	Description string         `gorm:"type:text" json:"description,omitempty"`

	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}


func (app *App) BeforeCreate(tx *gorm.DB) (err error) {
	app.ID = uuid.New()
	return
}

func (tenant *App) ToString() string {
	data, err := json.Marshal(tenant)
	if err != nil {
		return "error"
	}

	return string(data)
}
