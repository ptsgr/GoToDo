package service

import (
	"github.com/ptsgr/GoToDo"
	"github.com/ptsgr/GoToDo/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userID int, list GoToDo.TodoList) (int, error) {
	return s.repo.Create(userID, list)
}

func (s *TodoListService) GetAll(userID int) ([]GoToDo.TodoList, error) {
	return s.repo.GetAll(userID)
}
