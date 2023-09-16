package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucasmallmann/money-transfer/routes"
)

type Server struct {
	listenAddr string
	Client     *fiber.App
}

func NewServer(listenAddr string) *Server {
	app := fiber.New(fiber.Config{
		AppName: "money-transfer",
	})
	return &Server{listenAddr: listenAddr, Client: app}
}

func (s *Server) Serve() error {
	return s.Client.Listen(s.listenAddr)
}

func (s *Server) Setup() {
	routes.SetupRoutes(s.Client)
}
