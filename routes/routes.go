package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	router.GET("/events", getEvents)
	router.GET("/events/:id", getEventByID)
	router.POST("/events", createEvent)
	router.PUT("/events/:id", updateEventByID)
}
