package model

import (
	"time"

	"github.com/google/uuid"
)

type History struct {
	ID         uuid.UUID  `json:"id" db:"id"`
	UserID     uuid.UUID  `json:"user_id" db:"user_id"`
	BookID     uuid.UUID  `json:"book_id" db:"book_id"`
	BorrowedAt time.Time  `json:"borrowed_at" db:"borrowed_at"`
	ReturnedAt *time.Time `json:"returned_at" db:"returned_at"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
}
