package main

import (
	"Go_Exos/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Enregistrer les routes des tâches
	routes.RegisterTaskRoutes(r)

	// Démarrer le serveur
	r.Run(":8080")
}
