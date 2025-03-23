package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	InitRedis()

	r := gin.Default()

	// Apply our own metrics middleware
	r.Use(metricsMiddleware())

	// Register Prometheus metrics
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestDuration)

	// Routes
	r.POST("/validators", CreateValidatorRequest)
	r.GET("/validators/:request_id", CheckRequestStatus)

	// Expose /metrics for Prometheus
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Start server
	log.Println("Server running on port 8080")
	r.Run(":8080")
}
