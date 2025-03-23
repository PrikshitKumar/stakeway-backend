package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// Global variables
var (
	// Prometheus Metrics
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests received",
		},
		[]string{"method", "path"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
)

// Middleware to track metrics
func metricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.FullPath()

		log.Printf("Middleware triggered for path: %s", path)

		// Track request count
		requestCount.WithLabelValues(c.Request.Method, path).Inc()

		// Process request
		c.Next()

		// Track request duration
		duration := time.Since(start).Seconds()
		requestDuration.WithLabelValues(c.Request.Method, path).Observe(duration)
		log.Printf("Request completed: %s (Duration: %f seconds)", path, duration)
	}
}
