# Демо проект: ToDo APP, REST, Golang, SQLite

## Библиотеки:
* github.com/mattn/go-sqlite3
* github.com/gin-gonic/gin


## Запуск приложения
```
go run .\cmd\main.go
```
## Доступные пути(handlers)

* GET    /api/items/               
* PUT    /api/items/               
* GET    /api/items/:id            
* DELETE /api/items/:id        

## Использование / Тестирование

Приложение будет доступно по адресу http://localhost:8082

## Доступные пути(handlers)

* GET    /api/items/
* PUT    /api/items/
* GET    /api/items/:id
* DELETE /api/items/:id

Для тестов можно воспользоватся [Postman](docs/postman-test.md)

### Возможные ппроблемы.
* На Win для библиотеки github.com/mattn/go-sqlite3 потребовалось установить `gcc` compiler
* Более подробно в документации
 [https://github.com/mattn/go-sqlite3/blob/master/README.md#windows](https://github.com/mattn/go-sqlite3/blob/master/README.md#windows)