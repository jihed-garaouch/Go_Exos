package handlers

import (
	"GOExo/models"
	"net/http"
	"time"
    "sync"
	"log/slog"
	"github.com/gin-gonic/gin"
)

// Liste des tâches (stockée en mémoire)
var tasks = []models.Task{
	{ID: "1", Title: "Apprendre Go"},
	{ID: "2", Title: "Développer une API"},
}

// Récupérer toutes les tâches
func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}
func simulateTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	duration := time.Duration(id+2) * time.Second // Random duration between 2-5 seconds
	slog.Info("Task %d started, will take %v seconds\n", id, duration)
	time.Sleep(duration)
	slog.Info("Task %d completed\n", id)
}

func HandleMultipleTask(c *gin.Context) {
	var wg sync.WaitGroup
	numTasks := 5

	for i := 1; i <= numTasks; i++ {
		wg.Add(1)
		go simulateTask(i, &wg)
	}

	wg.Wait()
	 c.JSON(http.StatusOK, gin.H{"message": "All Task are Done"})

}

func PrintWithDelay(){
    time.Sleep(5 * time.Second)
    slog.Info("Task is Done")
}

func DelayProccessTask(c *gin.Context){

    go PrintWithDelay()
     c.JSON(http.StatusOK, gin.H{"message": "Task is Proccessing"})
}

// Ajouter une nouvelle tâche
func AddTask(c *gin.Context) {
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON invalide"})
		return
	}

	tasks = append(tasks, newTask)
	c.JSON(http.StatusCreated, newTask)
}

// Supprimer une tâche par ID
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	index := -1

	for i, task := range tasks {
		if task.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tâche non trouvée"})
		return
	}

	tasks = append(tasks[:index], tasks[index+1:]...)
	c.JSON(http.StatusOK, gin.H{"message": "Tâche supprimée avec succès"})
}
