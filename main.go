package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/lucasmallmann/money-transfer/api"
	"github.com/lucasmallmann/money-transfer/storage"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}
	postgresStorage, err := storage.NewPostgresStorage()

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(":3333", postgresStorage)
	server.Setup()
	log.Fatal(server.Serve())
}
