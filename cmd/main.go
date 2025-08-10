package main

import (
	"booking-api/db"
	"booking-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.SetupRoutes(server)
	server.Run(":8080") // Start the server on port 8080
}
