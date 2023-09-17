package main

import (
	"log"

	"github.com/lucasmallmann/money-transfer/api"
	"github.com/lucasmallmann/money-transfer/storage"
)

func main() {
	memoryStorage := storage.NewMemoryStorage()
	server := api.NewServer(":3333", memoryStorage)
	server.Setup()
	log.Fatal(server.Serve())
}
