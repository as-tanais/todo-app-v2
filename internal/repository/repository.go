package repository

import (
	"database/sql"
	"todo-app-v2/internal/model"
)

type TodoItem interface {
	Create(item model.TodoItem) (int64, error)
	GetAll() ([]model.TodoItem, error)
	Delete(itemId int) (int64, error)
	GetById(itemId int) (model.TodoItem, error)
}

type Repository struct {
	TodoItem
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		TodoItem: NewTodoItemSqlite(db),
	}
}
