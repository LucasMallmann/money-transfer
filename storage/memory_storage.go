package storage

import "github.com/lucasmallmann/money-transfer/types"

type MemoryStorage struct {
}

func (memory *MemoryStorage) GetByIds(id string) (*types.User, error) {
	return &types.User{
		Id:    "1",
		Name:  "Lucas",
		Email: "lucas@email.com",
	}, nil
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}
