package model

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	Title       string     `json:"title" db:"title"`
	Overview    string     `json:"overview" db:"overview"`
	ImageUrl    string     `json:"image_url" db:"image_url"`
	Isbn        *string    `json:"isbn" db:"isbn"`
	Writer      string     `json:"writer" db:"writer"`
	Publisher   string     `json:"publisher" db:"publisher"`
	PublishedAt int16      `json:"published_at" db:"published_at"`
	Category    string     `json:"category" db:"category"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
}
