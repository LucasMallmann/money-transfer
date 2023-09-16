package storage

import "github.com/lucasmallmann/money-transfer/types"

type UserStorage interface {
	GetById(id string) (*types.User, error)
}
