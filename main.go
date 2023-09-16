package main

import (
	"log"

	"github.com/lucasmallmann/money-transfer/api"
)

func main() {
	server := api.NewServer(":3333")
	server.Setup()
	log.Fatal(server.Serve())
}
