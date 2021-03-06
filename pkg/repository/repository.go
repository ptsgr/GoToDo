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
	GetAll(userID int) ([]GoToDo.TodoList, error)
	GetByID(userID, listID int) (GoToDo.TodoList, error)
	Delete(userID, listID int) error
	Update(userID, listID int, input GoToDo.UpdateListInput) error
}

type TodoItem interface {
	Create(listID int, item GoToDo.TodoItem) (int, error)
	GetAll(userID, listID int) ([]GoToDo.TodoItem, error)
	GetByID(userID, itemID int) (GoToDo.TodoItem, error)
	Delete(userID, itemID int) error
	Update(userID, itemID int, input GoToDo.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
