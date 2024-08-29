package main

import (
	todo_app_v2 "todo-app-v2"
	"todo-app-v2/pkg/handler"
	"todo-app-v2/pkg/repository"
	"todo-app-v2/pkg/service"
)

func main() {
	srv := new(todo_app_v2.Server)

	db := repository.NewSqliteDb("db.db")

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv.Run("8082", handlers.InitRoutes())
}
