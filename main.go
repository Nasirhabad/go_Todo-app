package main

import (
	"simple-to-do-app/models"
	"simple-to-do-app/services"
)

func main() {
	services.AddTodo(models.Todo)
	services.GetTodoList()
}
