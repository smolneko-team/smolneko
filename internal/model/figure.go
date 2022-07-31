package model

import (
    "time"

    "github.com/gofrs/uuid"
)

type Figure struct {
    ID           uuid.UUID `json:"id"` // uuid
    CharacterID  uuid.UUID `json:"character_id"`
    SourceID     uuid.UUID `json:"source_id"`
    Name         string    `json:"name"`
    Description  string    `json:"description"`
    Type         string    `json:"type"`
    Size         string    `json:"size"`
    Height       int       `json:"height"` // in millimeters
    Width        int       `json:"width"`  // in millimeters
    Materials    []string  `json:"material"`
    ReleaseDate  time.Time `json:"release_date"`
    Manufacturer string    `json:"manufacturer"`
    Link         string    `json:"link,omitempty"`
    Price        string    `json:"price"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
    DeletedAt    time.Time `json:"deleted_at,omitempty"`
    IsDraft      bool      `json:"is_draft"`
}

type Source struct {
    ID          uuid.UUID `json:"id"`
    Title       string    `json:"title"`
    ReleaseDate time.Time `json:"release_date"`
    Authors     []string  `json:"authors"`
}

type Character struct {
    ID      int       `json:"id"`
    Name    int       `json:"name"`
    TitleID uuid.UUID `json:"title_id"`
}

type Sculptor struct {
    ID       uuid.UUID `json:"id"`
    Name     string    `json:"name"`
    FigureID uuid.UUID `json:"figure_id"`
}