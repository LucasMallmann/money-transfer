package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucasmallmann/money-transfer/storage"
)

type Server struct {
	listenAddr string
	Client     *fiber.App
	storage    storage.Storage
}

func NewServer(listenAddr string, storage storage.Storage) *Server {
	app := fiber.New(fiber.Config{
		AppName: "money-transfer",
	})

	return &Server{
		listenAddr: listenAddr,
		Client:     app,
		storage:    storage,
	}
}

func (s *Server) Serve() error {
	return s.Client.Listen(s.listenAddr)
}

func (s *Server) Setup() {
	SetupRoutes(s.Client, s.storage)
}
