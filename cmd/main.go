package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	server.Run(":8080") // Start the server on port 8080
}
