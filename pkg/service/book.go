package service

import (
	"context"

	"application/pkg/dto"
	"application/pkg/model"
	"application/pkg/repository"

	"github.com/google/uuid"
)

type IBookService interface {
	GetAll(ctx context.Context, params dto.GetBooksRequest) ([]model.Book, model.Meta, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Book, error)
	Add(ctx context.Context, params dto.AddBookRequest) error
	Borrow(ctx context.Context, bookID uuid.UUID, userID uuid.UUID, params dto.BorrowBookRequest) error
	Update(ctx context.Context, bookID uuid.UUID, params dto.UpdateBookRequest) error
}

type BookService struct {
	Repositories repository.Repositories
}

func NewBook(repositories repository.Repositories) IBookService {
	return &BookService{Repositories: repositories}
}

func (s BookService) GetAll(ctx context.Context, params dto.GetBooksRequest) ([]model.Book, model.Meta, error) {
	return s.Repositories.Book.GetAll(ctx, params)
}

func (s BookService) GetByID(ctx context.Context, id uuid.UUID) (*model.Book, error) {
	return s.Repositories.Book.GetByID(ctx, id)
}

func (s BookService) Add(ctx context.Context, params dto.AddBookRequest) error {
	return s.Repositories.Book.Add(ctx, params)
}

func (s BookService) Borrow(ctx context.Context, bookID uuid.UUID, userID uuid.UUID, params dto.BorrowBookRequest) error {
	// TODO: check if user eligible to borrow book(s) or not
	// - Base on quantity of borrow
	// - and if there is any borrowed book off the limit
	_, err := s.Repositories.Book.GetByID(ctx, bookID)
	if err != nil {
		return err
	}
	// TODO: check book quantity - params.Quantity, less than 0 or not
	// TODO: update quantity of book
	// TODO: insert into history
	// TODO: update quantity of borrow user

	return nil
}

func (s BookService) Update(ctx context.Context, bookID uuid.UUID, params dto.UpdateBookRequest) error {
	return s.Repositories.Book.Update(ctx, bookID, params)
}
