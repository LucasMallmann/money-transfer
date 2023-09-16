package api

import "github.com/lucasmallmann/money-transfer/storage"

type UserStorage struct {
	storage storage.UserStorage
}

func NewUserStorage(storage storage.UserStorage) *UserStorage {
	return &UserStorage{storage: storage}
}
