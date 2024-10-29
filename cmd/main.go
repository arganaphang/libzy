package main

import (
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

	db, err := sqlx.Open(
		"postgres",
		os.Getenv("DATABASE_URL"),
	)
	if err != nil {
		log.Fatalln("failed to open database connection", err.Error())
	}
	if db.Ping() != nil {
		log.Fatalln("failed to ping database")
	}

	repositories := repository.Repositories{
		Book: repository.NewBook(db),
	}

	services := service.Services{
		Book: service.NewBook(repositories),
	}

	// Unused return handlers, just for grouping
	_ = handler.Handlers{
		Book: handler.NewBook(app, services),
	}

	app.Get("/healthz", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(map[string]string{
			"message": "OK",
		})
	})

	app.Listen("0.0.0.0:8000")
}
