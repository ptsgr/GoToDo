package service

import (
	"github.com/ptsgr/GoToDo"
	"github.com/ptsgr/GoToDo/pkg/repository"
)

type Authorization interface {
	CreateUser(user GoToDo.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
