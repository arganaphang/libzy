package model

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	UserRoleLibrarian UserRole = "librarian"
	UserRoleMember    UserRole = "member"
)

type User struct {
	ID               uuid.UUID  `json:"id" db:"id"`
	Name             string     `json:"name" db:"name"`
	Email            string     `json:"email" db:"email"`
	Username         string     `json:"username" db:"username"`
	Password         string     `json:"password" db:"password"`
	Role             UserRole   `json:"role" db:"role"`
	QuantityOfBorrow int16      `json:"quantity_of_borrow" db:"quantity_of_borrow"`
	Address          string     `json:"address" db:"address"`
	ImageUrl         *string    `json:"image_url" db:"image_url"`
	CreatedAt        time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at" db:"deleted_at"`
}
