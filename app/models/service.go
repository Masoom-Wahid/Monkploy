package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	ID        uuid.UUID        `gorm:"primaryKey" json:"id,omitempty"`
	Name        string         `gorm:"not null" json:"name,omitempty"`
	Description string         `gorm:"type:text" json:"description,omitempty"`
	Type        string    	   `gorm:"type:varchar(50);not null" json:"type,omitempty"`



	// Relations
	AppID  uuid.UUID 

	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}


func (s *Service) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}

func (s *Service) ToString() string {
	data, err := json.Marshal(s)
	if err != nil {
		return "error"
	}

	return string(data)
}
