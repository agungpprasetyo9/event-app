package main

import (
	"log"

	"example.com/event-app/config"
	"example.com/event-app/controllers"
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
		api.GET("/event", controllers.GetEvents)
		api.POST("/event", controllers.CreateEvent)
		api.GET("/event/:id", controllers.GetEventsbyId)
		api.PUT("/event/:id", controllers.UpdateEvent)
		api.DELETE("/event/:id", controllers.DeleteEvent)

		api.POST("/auth/register", controllers.RegisterUser)
		api.POST("/auth/login", controllers.LoginUser)
	}

	server.Run(":8080")
}
