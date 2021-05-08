package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ptsgr/GoToDo"
)

type Authorization interface {
	CreateUser(user GoToDo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
