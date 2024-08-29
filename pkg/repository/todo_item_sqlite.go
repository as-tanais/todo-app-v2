package repository

import (
	"database/sql"
	"fmt"
	"todo-app-v2/pkg/model"
)

type TodoItemSqlite struct {
	db *sql.DB
}

func NewTodoItemSqlite(db *sql.DB) *TodoItemSqlite {
	return &TodoItemSqlite{db: db}
}

func (r *TodoItemSqlite) Delete(itemId int) (int64, error) {
	sqlQuery := "DELETE FROM tasks WHERE id = ?"

	stmt, err := r.db.Prepare(sqlQuery)

	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(itemId)

	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}

func (r *TodoItemSqlite) Create(item model.TodoItem) (int64, error) {
	sqlQuery := "INSERT INTO tasks(title, description, done) VALUES(?, ?, ?)"

	stmt, err := r.db.Prepare(sqlQuery)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	res, err2 := stmt.Exec(item.Title, item.Description, false)

	if err2 != nil {
		panic(err2)
	}
	return res.LastInsertId()

}

func (r *TodoItemSqlite) GetAll() ([]model.TodoItem, error) {
	var items []model.TodoItem
	query := fmt.Sprintf(`SELECT * FROM tasks`)

	rows, err := r.db.Query(query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		todo := model.TodoItem{}
		err2 := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Done)
		if err2 != nil {
			panic(err2)
		}
		items = append(items, todo)
	}

	return items, nil
}
