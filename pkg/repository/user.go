package repository

import (
	"context"

	"application/pkg/dto"
	"application/pkg/model"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type IUserRepository interface {
	GetAll(ctx context.Context, params dto.GetUsersRequest) ([]model.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	Register(ctx context.Context, params dto.RegisterRequest) error
	Update(ctx context.Context, id uuid.UUID, params dto.UpdateUserRequest) error
}

type UserRepository struct {
	DB *sqlx.DB
}

func NewUser(db *sqlx.DB) IUserRepository {
	return &UserRepository{DB: db}
}

func (r UserRepository) GetAll(ctx context.Context, params dto.GetUsersRequest) ([]model.User, error) {
	return nil, nil
}

func (r UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return nil, nil
}

func (r UserRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	return nil, nil
}

func (r UserRepository) Register(ctx context.Context, params dto.RegisterRequest) error {
	return nil
}

func (r UserRepository) Update(ctx context.Context, id uuid.UUID, params dto.UpdateUserRequest) error {
	return nil
}
