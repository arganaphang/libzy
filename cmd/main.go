package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"application/pkg/handler"
	"application/pkg/repository"
	"application/pkg/service"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()
	app.Get("/healthz", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(map[string]string{
			"message": "OK",
		})
	})

	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf(
			"postgres://%s:%s@%s:5432/%s?sslmode=disable",
			os.Getenv("DATABASE_USER"),
			os.Getenv("DATABASE_PASS"),
			os.Getenv("DATABASE_HOST"),
			os.Getenv("DATABASE_NAME"),
		),
	)
	if err != nil {
		log.Fatalln("failed to open database connection", err.Error())
	}
	if db.Ping() != nil {
		log.Fatalln("failed to ping database")
	}

	repositories := repository.Repositories{
		Book:    repository.NewBook(db),
		User:    repository.NewUser(db),
		History: repository.NewHistory(db),
	}

	services := service.Services{
		Book: service.NewBook(repositories),
		User: service.NewUser(repositories),
	}

	// Unused return handlers, just for grouping
	_ = handler.Handlers{
		Book: handler.NewBook(app, services),
		User: handler.NewUser(app, services),
	}

	app.Listen("0.0.0.0:8000")
}
