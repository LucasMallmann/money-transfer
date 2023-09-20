package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/lucasmallmann/money-transfer/config"
	"github.com/lucasmallmann/money-transfer/types"
)

type Storage interface {
	GetById(id string) (*types.User, error)
}

type PostgresStorage struct {
	db *sqlx.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	driver, connString := config.NewDbConfig("postgres").GetConnString()
	db, err := sqlx.Connect(driver, connString)
	if err != nil {
		return nil, err
	}
	return &PostgresStorage{
		db: db,
	}, nil
}

func (postgres *PostgresStorage) GetById(id string) (*types.User, error) {
	return &types.User{
		Id:    "1",
		Name:  "Lucas",
		Email: "lucas@email.com",
	}, nil
}
