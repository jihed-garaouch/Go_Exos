package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var tasks = []Task{
	{ID: 1, Title: "Task 1"},
	{ID: 2, Title: "Task 2"},
}
var nextID = 3

func main() {
	r := gin.Default()

	r.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, tasks)
	})
	r.DELETE("/tasks/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
			return
		}
		for i, task := range tasks {
			if id == task.ID {
				tasks = append(tasks[:i], tasks[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "task est supprimé"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "task n'est pas trouvé"})
	})

	r.Run(":8080")
}
