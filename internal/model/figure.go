package model

import (
	"time"
)

type Figure struct {
	ID           string     `json:"id"`
	CharacterID  *int       `json:"character_id,omitempty"`
	Name         string     `json:"name"`
	PreviewImage string     `json:"preview_image"`
	Description  *string    `json:"description,omitempty"`
	Type         *string    `json:"type,omitempty"`
	Version      *string    `json:"version"`
	Size         *string    `json:"size,omitempty"`
	Height       *int       `json:"height,omitempty"`
	Materials    *[]string  `json:"materials,omitempty"`
	ReleaseDate  *time.Time `json:"release_date,omitempty"`
	Manufacturer *string    `json:"manufacturer,omitempty"`
	Links        *[]string  `json:"links,omitempty"`
	Price        *[]string  `json:"price,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
	IsNSFW       bool       `json:"is_nsfw"`
	IsDraft      bool       `json:"is_draft"`
}
