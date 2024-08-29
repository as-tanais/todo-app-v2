package service

import (
	"todo-app-v2/pkg/model"
	"todo-app-v2/pkg/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
}

func (s *TodoItemService) Delete(itemId int) (int64, error) {
	return s.repo.Delete(itemId)
}

func (s *TodoItemService) Create(item model.TodoItem) (int64, error) {
	return s.repo.Create(item)
}

func NewTodoItemService(repo repository.TodoItem) *TodoItemService {
	return &TodoItemService{repo: repo}
}

func (s *TodoItemService) GetAll() ([]model.TodoItem, error) {
	return s.repo.GetAll()
}
