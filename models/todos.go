package models

import "time"

type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func GetAllTodos(todos *[]Todo) {
	Db.Find(&todos)
}

func CreateTodo(todo *Todo) {
	Db.NewRecord(todo)
	Db.Create(&todo)
}

func DeleteTodo(todo *Todo, id string) {
	Db.Delete(&todo, id)
}
