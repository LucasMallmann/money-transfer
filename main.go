package main

import (
	"github.com/lucasmallmann/money-transfer/app"
)

func main() {
	app := app.NewApp("money transfer")
	app.Setup()
	app.Serve(":3333")
}
