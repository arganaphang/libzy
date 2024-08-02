package repository

import (
	"context"

	"application/pkg/dto"
	"application/pkg/model"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type IHistoryRepository interface {
	GetAllByUserID(ctx context.Context, userID uuid.UUID) ([]model.History, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.History, error)
	Add(ctx context.Context, params dto.AddHistoryRequest) error
	Update(ctx context.Context, id uuid.UUID, params dto.UpdateHistoryRequest) error
}

type HistoryRepository struct {
	DB *sqlx.DB
}

func NewHistory(db *sqlx.DB) IHistoryRepository {
	return &HistoryRepository{DB: db}
}

func (r HistoryRepository) GetAllByUserID(ctx context.Context, userID uuid.UUID) ([]model.History, error) {
	return nil, nil
}

func (r HistoryRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.History, error) {
	return nil, nil
}

func (r HistoryRepository) Add(ctx context.Context, params dto.AddHistoryRequest) error {
	return nil
}

func (r HistoryRepository) Update(ctx context.Context, id uuid.UUID, params dto.UpdateHistoryRequest) error {
	return nil
}
