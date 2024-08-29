package main

import (
	todo_app_v2 "todo-app-v2"
	"todo-app-v2/pkg/handler"
	"todo-app-v2/pkg/repository"
	"todo-app-v2/pkg/service"
)

func main() {
	//создаем сервер
	srv := new(todo_app_v2.Server)
	//создаем базу
	db := repository.NewSqliteDb("db.db")
	//создаем все зависимости последовательно
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	//запускаем сервер на порту 8082
	srv.Run("8082", handlers.InitRoutes())
}
