package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ptsgr/GoToDo"
)

type Authorization interface {
	CreateUser(user GoToDo.User) (int, error)
	GetUser(username, password string) (GoToDo.User, error)
}

type TodoList interface {
	Create(userID int, list GoToDo.TodoList) (int, error)
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
		TodoList:      NewTodoLIstPostgres(db),
	}
}
