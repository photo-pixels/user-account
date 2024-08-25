package storage

import (
	"errors"
)

var (
	// ErrNotFound ошибка не найденно
	ErrNotFound = errors.New("not found")
	// ErrAlreadyExist запись уже существует
	ErrAlreadyExist = errors.New("already exist")
)
