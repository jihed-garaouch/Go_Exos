package routes

import (
	"GOExo/handlers"

	"github.com/gin-gonic/gin"
)

// Configurer les routes des tâches
func RegisterTaskRoutes(router *gin.Engine) {
	router.GET("/tasks", handlers.GetTasks)          // Récupérer toutes les tâches
	router.POST("/tasks", handlers.AddTask)          // Ajouter une nouvelle tâche
	router.DELETE("/tasks/:id", handlers.DeleteTask) // Supprimer une tâche
	router.GET("/tasks/process", handlers.DelayProccessTask)

}
