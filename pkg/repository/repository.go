package repository

import (
	"database/sql"
	"todo-app-v2/pkg/model"
)

type TodoItem interface {
	Create(item model.TodoItem) (int64, error)
	GetAll() ([]model.TodoItem, error)
	Delete(itemId int) (int64, error)
}

type Repository struct {
	TodoItem
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		TodoItem: NewTodoItemSqlite(db),
	}
}
