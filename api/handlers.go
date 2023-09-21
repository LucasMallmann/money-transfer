package api

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/lucasmallmann/money-transfer/storage"
	"github.com/lucasmallmann/money-transfer/types"
)

type UserHandler struct {
	storage storage.Storage
}

type TransferHandler struct {
	storage storage.Storage
}

func NewUserHandler(storage storage.Storage) *UserHandler {
	return &UserHandler{
		storage: storage,
	}
}

func NewTransferHandler(storage storage.Storage) *TransferHandler {
	return &TransferHandler{
		storage: storage,
	}
}

func (handler *TransferHandler) HandleTransfer(ctx *fiber.Ctx) error {
	transferRequest := &types.TransferRequest{}
	ctx.BodyParser(transferRequest)

	validator := validator.New()

	err := validator.Struct(transferRequest)

	if err != nil {
		log.Println(err)
		return ctx.Status(400).SendString(fmt.Sprintf("Not valid request %s\n", err.Error()))
	}

	hasMoney, err := handler.storage.CheckBalance(transferRequest.Sender, transferRequest.Amount)

	if err != nil {
		log.Println(err)
		return ctx.Status(500).SendString("Error checking balance")
	}

	if !hasMoney {
		log.Println(err)
		return ctx.Status(400).JSON(fiber.Map{
			"error":   "InsufficientFunds",
			"message": "Error checking balance",
		})
	}

	// Check if user has enough balance
	err = handler.storage.
		TransferAmount(
			transferRequest.Sender,
			transferRequest.Receiver,
			transferRequest.Amount,
		)

	if err != nil {
		log.Println(err)
		return ctx.Status(500).SendString("Error transferring amount")
	}

	fmt.Printf("user has enough money: %v\n", hasMoney)

	return ctx.SendString("Transfer done!")
}
