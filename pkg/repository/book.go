package repository

import (
	"context"

	"application/pkg/dto"
	"application/pkg/model"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/doug-martin/goqu/v9"
)

type IBookRepository interface {
	GetAll(ctx context.Context, params dto.GetBooksRequest) ([]model.Book, model.Meta, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Book, error)
	Add(ctx context.Context, params dto.AddBookRequest) error
	Update(ctx context.Context, id uuid.UUID, params dto.UpdateBookRequest) error
}

type BookRepository struct {
	DB *sqlx.DB
}

func NewBook(db *sqlx.DB) IBookRepository {
	return &BookRepository{DB: db}
}

func (r BookRepository) GetAll(ctx context.Context, params dto.GetBooksRequest) ([]model.Book, model.Meta, error) {
	meta := model.Meta{
		Page:    params.Page,
		PerPage: params.PerPage,
		Total:   0,
	}

	offset := (params.Page - 1) * params.PerPage

	query := goqu.From("books").Limit(params.PerPage).Offset(offset)

	sql, _, err := query.ToSQL()
	if err != nil {
		return nil, meta, err
	}

	var results []model.Book
	if err := r.DB.Select(&results, sql); err != nil {
		return nil, meta, err
	}

	sqlCount, _, err := query.Select(goqu.COUNT("*")).ToSQL()
	if err != nil {
		return nil, meta, err
	}

	var count int
	if err := r.DB.Get(&count, sqlCount); err != nil {
		return nil, meta, err
	}

	meta.Total = uint(count)

	return results, meta, nil
}

func (r BookRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Book, error) {
	sql, _, err := goqu.From("books").Where(goqu.C("id").Eq(id)).Limit(1).ToSQL()
	if err != nil {
		return nil, err
	}

	var result model.Book
	if err := r.DB.Get(&result, sql); err != nil {
		logrus.Info("Error ", err.Error())
		return nil, err
	}

	return &result, nil
}

func (r BookRepository) Add(ctx context.Context, params dto.AddBookRequest) error {
	sql, _, err := goqu.
		Insert("books").
		Cols(
			"title",
			"overview",
			"image_url",
			"isbn",
			"writer",
			"publisher",
			"published_at",
			"category",
		).
		Vals(goqu.Vals{
			params.Title,
			params.Overview,
			params.ImageUrl,
			params.Isbn,
			params.Writer,
			params.Publisher,
			params.PublishedAt,
			params.Category,
		}).
		ToSQL()
	if err != nil {
		return err
	}
	if _, err := r.DB.Exec(sql); err != nil {
		return err
	}
	return nil
}

func (r BookRepository) Update(ctx context.Context, id uuid.UUID, params dto.UpdateBookRequest) error {
	sql, _, err := goqu.
		Update("books").
		Set(
			goqu.Record{
				"title":        params.Title,
				"overview":     params.Overview,
				"image_url":    params.ImageUrl,
				"isbn":         params.Isbn,
				"writer":       params.Writer,
				"publisher":    params.Publisher,
				"published_at": params.PublishedAt,
				"category":     params.Category,
			}).
		Where(goqu.C("id").Eq(id)).
		ToSQL()
	if err != nil {
		return err
	}
	if _, err := r.DB.Exec(sql); err != nil {
		return err
	}
	return nil
}
