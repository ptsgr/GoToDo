package service

import (
	"github.com/ptsgr/GoToDo"
	"github.com/ptsgr/GoToDo/pkg/repository"
)

type Authorization interface {
	CreateUser(user GoToDo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userID int, list GoToDo.TodoList) (int, error)
	GetAll(userID int) ([]GoToDo.TodoList, error)
	GetByID(userID, listID int) (GoToDo.TodoList, error)
}

type TodoItem interface {
	Create(userID, listID int, item GoToDo.TodoItem) (int, error)
	GetAll(userID, listID int) ([]GoToDo.TodoItem, error)
	GetByID(userID, itemID int) (GoToDo.TodoItem, error)
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
