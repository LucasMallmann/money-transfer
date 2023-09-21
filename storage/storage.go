package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/lucasmallmann/money-transfer/config"
	"github.com/lucasmallmann/money-transfer/types"
)

type Storage interface {
	TransferAmount(senderId string, receiver string, amount float64) error
	CheckBalance(id string, amount float64) (bool, error)
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

func (postgres *PostgresStorage) CheckBalance(id string, amount float64) (bool, error) {
	query := `
    SELECT CASE
      WHEN u.balance >= $1 THEN true
      ELSE false
    END AS transfer_status
    FROM users u
    WHERE u.id = $2
  `
	var transferStatus bool
	err := postgres.db.Get(&transferStatus, query, amount, id)
	if err != nil {
		return false, err
	}
	return transferStatus, nil
}

func (p *PostgresStorage) TransferAmount(senderId string, receiver string, amount float64) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var sender types.User
	err = p.db.Get(&sender, "SELECT * FROM users WHERE id = $1", senderId)

	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE users SET balance = balance - $1 WHERE id = $2", amount, senderId)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Update receiver balance
	_, err = tx.Exec("UPDATE users SET balance = balance + $1 WHERE id = $2", amount, receiver)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert into transfers table
	_, err = tx.Exec(`
    INSERT INTO transferss (id_sender, id_receiver, value) 
    VALUES ($1, $2, $3)`, senderId, receiver, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
