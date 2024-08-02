package handler

import (
	"net/http"

	"application/pkg/dto"
	"application/pkg/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type IBookHandler interface {
	GetAll(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
	Add(ctx *fiber.Ctx) error
	Borrow(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}

type BookHandler struct {
	Services service.Services
}

func NewBook(app *fiber.App, services service.Services) IBookHandler {
	handler := &BookHandler{Services: services}

	route := app.Group("/book")
	route.Get("/", handler.GetAll)
	route.Post("/", handler.Add)
	route.Get("/:id", handler.GetByID)
	route.Put("/:id", handler.Update)
	route.Put("/:id/borrow", handler.Borrow)

	return handler
}

func (h BookHandler) GetAll(ctx *fiber.Ctx) error {
	var params dto.GetBooksRequest
	if err := ctx.QueryParser(&params); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.GetBookByIDResponse{
			Success: false,
			Message: "failed to parse query params",
		})
	}
	if params.Page < 1 {
		params.Page = 1
	}
	result, meta, err := h.Services.Book.GetAll(ctx.UserContext(), params)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.GetBooksResponse{
			Success: false,
			Message: "failed to get all",
		})
	}
	return ctx.Status(http.StatusOK).JSON(dto.GetBooksResponse{
		Success: true,
		Message: "get all book successfully",
		Data:    result,
		Meta:    meta,
	})
}

func (h BookHandler) GetByID(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.GetBookByIDResponse{
			Success: false,
			Message: "id must be uuid",
		})
	}
	result, err := h.Services.Book.GetByID(ctx.UserContext(), id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.GetBookByIDResponse{
			Success: false,
			Message: "failed to get by id",
		})
	}
	return ctx.Status(http.StatusOK).JSON(dto.GetBookByIDResponse{
		Success: true,
		Message: "get book by id successfully",
		Data:    result,
	})
}

func (h BookHandler) Add(ctx *fiber.Ctx) error {
	var params dto.AddBookRequest
	if err := ctx.BodyParser(&params); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.AddBookResponse{
			Success: false,
			Message: "failed to parse body",
		})
	}
	if err := h.Services.Book.Add(ctx.UserContext(), params); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.AddBookResponse{
			Success: false,
			Message: "failed to add book",
		})
	}
	return ctx.Status(http.StatusOK).JSON(dto.AddBookResponse{
		Success: true,
		Message: "add book success",
	})
}

func (h BookHandler) Borrow(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.GetBookByIDResponse{
			Success: false,
			Message: "id must be uuid",
		})
	}
	var params dto.BorrowBookRequest
	if err := ctx.BodyParser(&params); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.BorrowBookResponse{
			Success: false,
			Message: "failed to parse body",
		})
	}
	// TODO: Get User ID
	if err := h.Services.Book.Borrow(ctx.UserContext(), id, uuid.Nil, params); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.BorrowBookResponse{
			Success: false,
			Message: "failed to borrow book",
		})
	}
	return ctx.Status(http.StatusOK).JSON(dto.BorrowBookResponse{
		Success: true,
		Message: "borrow book success",
	})
}

func (h BookHandler) Update(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.GetBookByIDResponse{
			Success: false,
			Message: "id must be uuid",
		})
	}
	var params dto.UpdateBookRequest
	if err := ctx.BodyParser(&params); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.UpdateBookResponse{
			Success: false,
			Message: "failed to parse body",
		})
	}
	if err := h.Services.Book.Update(ctx.UserContext(), id, params); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.UpdateBookResponse{
			Success: false,
			Message: "failed to update book",
		})
	}
	return ctx.Status(http.StatusOK).JSON(dto.UpdateBookResponse{
		Success: true,
		Message: "update book success",
	})
}
