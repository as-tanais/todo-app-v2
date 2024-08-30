package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteDb(filepath string) *sql.DB {
	//откроем файл или создадим его
	db, err := sql.Open("sqlite3", filepath)

	// проверяем ошибки и выходим при их наличии
	if err != nil {
		panic(err)
	}

	// если ошибок нет, но не можем подключиться к базе данных,
	// то так же выходим
	if db == nil {
		panic("db nil")
	}

	migrate(db)

	return db
}

func migrate(db *sql.DB) {
	sqlQuery := `
    CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        title VARCHAR NOT NULL,
        description VARCHAR,
        done boolean                            
    );
    `

	_, err := db.Exec(sqlQuery)
	// выходим, если будут ошибки с SQL запросом выше
	if err != nil {
		panic(err)
	}
}
