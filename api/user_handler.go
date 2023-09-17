package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucasmallmann/money-transfer/storage"
)

type UserHandler struct {
	storage storage.UserStorage
}

func NewUserHandler(storage storage.UserStorage) *UserHandler {
	return &UserHandler{
		storage: storage,
	}
}

func (h *UserHandler) Index(ctx *fiber.Ctx) error {
	user, err := h.storage.GetById("1")
	if err != nil {
		return ctx.SendStatus(500)
	}
	return ctx.Status(200).JSON(user)
}
