package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// Struct for request body
type RequestBody struct {
	NumValidators int    `json:"num_validators"`
	FeeRecipient  string `json:"fee_recipient"`
}

// Create Validator Request API
func CreateValidatorRequest(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	requestID := uuid.NewV4().String()
	validatorRequest := ValidatorRequest{
		RequestID:     requestID,
		NumValidators: body.NumValidators,
		FeeRecipient:  body.FeeRecipient,
		Status:        "started",
	}

	// Store request in Redis (as JSON)
	data, err := json.Marshal(validatorRequest)
	if err != nil {
		log.Println("Error marshaling request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process request"})
		return
	}

	err = rdb.Set(ctx, requestID, data, 24*time.Hour).Err()
	if err != nil {
		log.Println("Error saving request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process request"})
		return
	}

	// Start async key generation
	go ProcessValidatorRequest(requestID, body.NumValidators)

	// Respond with request ID
	c.JSON(http.StatusOK, gin.H{
		"request_id": requestID,
		"message":    "Validator creation in progress",
	})
}

// Check Request Status API
func CheckRequestStatus(c *gin.Context) {
	requestID := c.Param("request_id")

	validatorRequest, err := GetRequestStatus(requestID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
		return
	}

	c.JSON(http.StatusOK, validatorRequest)
}
