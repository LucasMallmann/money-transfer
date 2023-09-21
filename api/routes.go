package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucasmallmann/money-transfer/storage"
)

func SetupRoutes(app *fiber.App, storage storage.Storage) {
	transferHandler := NewTransferHandler(storage)

	app.Post("/transfer", transferHandler.HandleTransfer)
}
