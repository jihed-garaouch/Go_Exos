package models

// Task représente une tâche avec un ID et un titre
type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
