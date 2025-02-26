package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var dictionary = make(map[string]string)

func main() {
	r := gin.Default()

	r.GET("/words", func(c *gin.Context) {
		c.JSON(http.StatusOK, dictionary)
	})

	r.POST("/add", func(c *gin.Context) {
		var request struct {
			Word       string `json:"word"`
			Definition string `json:"definition"`
		}

		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON invalide"})
			return
		}

		dictionary[request.Word] = request.Definition

		c.JSON(http.StatusOK, gin.H{"message": "Mot ajouté avec succès"})
	})

	r.Run(":8080")
}
