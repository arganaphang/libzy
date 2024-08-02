package handler

import (
	"net/http"

	"application/pkg/service"

	"github.com/gofiber/fiber/v2"
)

type IUserHandler interface {
	GetAll(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}

type UserHandler struct {
	Services service.Services
}

func NewUser(app *fiber.App, services service.Services) IUserHandler {
	handler := &UserHandler{Services: services}

	route := app.Group("/user")
	route.Get("/", handler.GetAll)
	route.Get("/:id", handler.GetByID)
	route.Post("/login", handler.Login)
	route.Post("/register", handler.Register)

	return handler
}

func (h UserHandler) Login(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusInternalServerError).JSON(map[string]string{
		"message": "unimplemented",
	})
}

func (h UserHandler) Register(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusInternalServerError).JSON(map[string]string{
		"message": "unimplemented",
	})
}

func (h UserHandler) GetByID(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusInternalServerError).JSON(map[string]string{
		"message": "unimplemented",
	})
}

func (h UserHandler) GetAll(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusInternalServerError).JSON(map[string]string{
		"message": "unimplemented",
	})
}
