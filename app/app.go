package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lucasmallmann/money-transfer/routes"
)

type App struct {
	AppName string
	Client  *fiber.App
}

func NewApp(name string) App {
	app := fiber.New(fiber.Config{
		AppName: name,
	})

	return App{
		AppName: name,
		Client:  app,
	}
}

func (a *App) Serve(addr string) {
	log.Fatal(a.Client.Listen(addr))
}

func (a *App) Setup() {
	routes.SetupRoutes(a.Client)
}
