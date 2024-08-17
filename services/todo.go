package services

import (
	"fmt"
	"simple-to-do-app/models"
)

var database = []models.Todo = []models.Todo{}

func main() {
	// specify struct / function with uppercase letter first -> accessible across packages 
	// ec: Todo, AddtoDB()
	//if we specify  struct of function with lowercase first it wil be private (only accesseble one)

}

func menu() {
	fmt.Println("")
}

func AddTodo(new models.Todo) {
	database = append(database, newTodo)
}

func GetTodoList() []models.Todo {
	return database
}
