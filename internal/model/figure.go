package model

import (
	"time"
)

type Figure struct {
	ID           int        `json:"id"`
	CharacterID  *int       `json:"character_id,omitempty"`
	Name         string     `json:"name"`
	Description  *string    `json:"description,omitempty"`
	Type         *string    `json:"type,omitempty"`
	Size         *string    `json:"size,omitempty"`
	Height       *int       `json:"height,omitempty"` // in millimeters
	Materials    *[]string  `json:"materials,omitempty"`
	ReleaseDate  *time.Time `json:"release_date,omitempty"`
	Manufacturer *string    `json:"manufacturer,omitempty"`
	Links        *[]string  `json:"links,omitempty"`
	Price        *[]string  `json:"price,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
	IsDraft      bool       `json:"is_draft"`
}
