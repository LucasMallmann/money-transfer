package storage

import "github.com/lucasmallmann/money-transfer/types"

type GetById interface {
	GetById(id string) (*types.User, error)
}

type UserStorage interface {
	GetById
}
