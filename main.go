package main

import (
	"log"
	"net/http"

	"example.com/event-app/config"
	"example.com/event-app/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()
	server := gin.Default()

	// route
	api := server.Group("/api")
	{
		api.GET("/event", getEvent)
		api.POST("/event", createEvent)
	}

	server.Run(":8080")
}

// test func
func showTest(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "ngetest doang.....",
	})
}

// function Handle
func getEvent(context *gin.Context) {
	events := models.GetAllEvent()

	context.JSON(http.StatusOK, events)

}

// buat event
func createEvent(context *gin.Context) {

	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
			"error":   err.Error(),
		})
		return
	}
	event.Name = "Agung"
	event.Location = "Timika"
	event.Description = "growth people"

	// save
	event.Save()

	context.JSON(http.StatusCreated, gin.H{
		"Message": "event created",
		"event":   event,
	})
}
