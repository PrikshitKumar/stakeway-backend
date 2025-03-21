package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	InitRedis()

	r := gin.Default()

	// Routes
	r.POST("/validators", CreateValidatorRequest)
	r.GET("/validators/:request_id", CheckRequestStatus)

	// Start server
	log.Println("Server running on port 8080")
	r.Run(":8080")
}
