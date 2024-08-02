package dto

import (
	"application/pkg/model"
)

type GetBooksRequest struct {
	Q       *string `query:"q"`
	Page    uint    `query:"page"`
	PerPage uint    `query:"per_page"`
}

type GetBooksResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    []model.Book `json:"data"`
	Meta    model.Meta   `json:"meta"`
}

type AddBookRequest struct {
	Title       string  `json:"title"`
	Overview    string  `json:"overview"`
	Isbn        *string `json:"isbn"`
	Writer      string  `json:"writer"`
	Publisher   string  `json:"publisher"`
	PublishedAt uint16  `json:"published_at"`
	Category    string  `json:"category"`
	ImageUrl    string  `json:"image_url"`
}

type AddBookResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type GetBookByIDResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    *model.Book `json:"data"`
}

type BorrowBookRequest struct {
	Quantity uint8 `json:"quantity"`
}

type BorrowBookResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type UpdateBookRequest struct {
	Title       string  `json:"title"`
	Overview    string  `json:"overview"`
	Isbn        *string `json:"isbn"`
	Writer      string  `json:"writer"`
	Publisher   string  `json:"publisher"`
	PublishedAt uint16  `json:"published_at"`
	Category    string  `json:"category"`
	ImageUrl    string  `json:"image_url"`
}

type UpdateBookResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
