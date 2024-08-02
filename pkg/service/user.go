package service

import (
	"context"

	"application/pkg/dto"
	"application/pkg/model"
	"application/pkg/repository"

	"github.com/google/uuid"
)

type IUserService interface {
	GetAll(ctx context.Context, params dto.GetUsersRequest) ([]model.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	Login(ctx context.Context, params dto.LoginRequest) error
	Register(ctx context.Context, params dto.RegisterRequest) error
}

type UserService struct {
	Repositories repository.Repositories
}

func NewUser(repositories repository.Repositories) IUserService {
	return &UserService{Repositories: repositories}
}

func (s UserService) GetAll(ctx context.Context, params dto.GetUsersRequest) ([]model.User, error) {
	return s.Repositories.User.GetAll(ctx, params)
}

func (s UserService) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return s.Repositories.User.GetByID(ctx, id)
}

func (s UserService) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	return s.Repositories.User.GetByUsername(ctx, username)
}

func (s UserService) Login(ctx context.Context, params dto.LoginRequest) error {
	// TODO: implement this
	return nil
}

func (s UserService) Register(ctx context.Context, params dto.RegisterRequest) error {
	return s.Repositories.User.Register(ctx, params)
}
