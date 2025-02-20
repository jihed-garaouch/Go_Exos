package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	r.GET("/todo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"id":   0,
			"task": "firstTasks",
		})
	})

	r.Run(":8080")
}
