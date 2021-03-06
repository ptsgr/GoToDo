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
	Delete(userID, listID int) error
	Update(userID, listID int, input GoToDo.UpdateListInput) error
}

type TodoItem interface {
	Create(userID, listID int, item GoToDo.TodoItem) (int, error)
	GetAll(userID, listID int) ([]GoToDo.TodoItem, error)
	GetByID(userID, itemID int) (GoToDo.TodoItem, error)
	Delete(userID, itemID int) error
	Update(userID, itemID int, input GoToDo.UpdateItemInput) error
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
