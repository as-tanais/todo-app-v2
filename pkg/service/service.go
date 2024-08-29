package service

import (
	"todo-app-v2/pkg/model"
	"todo-app-v2/pkg/repository"
)

type TodoItem interface {
	GetAll() ([]model.TodoItem, error)
	Create(item model.TodoItem) (int64, error)
	Delete(itemId int) (int64, error)
}

type Service struct {
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TodoItem: NewTodoItemService(repos.TodoItem),
	}
}
