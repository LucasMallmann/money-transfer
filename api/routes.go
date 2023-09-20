package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lucasmallmann/money-transfer/storage"
)

func SetupRoutes(app *fiber.App, storage storage.Storage) {
	fmt.Println("Setting up routes...")
	userHandler := NewUserHandler(storage)
	app.Get("/users", userHandler.Index)
}
