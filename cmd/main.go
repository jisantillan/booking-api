package main

import (
	"booking-api/db"
	"booking-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") // Start the server on port 8080
}

func getEvents(context *gin.Context) {
	events, err := model.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event model.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not create event"})
		return
	}

	createdEvent := model.Save(event)

	context.JSON(http.StatusCreated, createdEvent)

}
