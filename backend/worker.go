package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Simulate validator key generation
func GenerateValidatorKeys(num int) []string {
	keys := make([]string, num)
	for i := 0; i < num; i++ {
		log.Printf("Generating key %d/%d... (waiting 20ms)", i+1, num)
		time.Sleep(20 * time.Millisecond) // Simulate processing time
		keys[i] = fmt.Sprintf("key_%d_%d", i+1, rand.Intn(10000))
		log.Printf("Key %d generated: %s", i+1, keys[i])
	}
	return keys
}

// Process validator request asynchronously
func ProcessValidatorRequest(requestID string, numValidators int) {
	log.Printf("Processing validator request: %s", requestID)

	keys := GenerateValidatorKeys(numValidators)
	if len(keys) == 0 {
		log.Printf("Failed to generate keys for request: %s", requestID)
		SetRequestStatus(requestID, "failed", nil)
		return
	}

	// Successfully generated keys, update status
	log.Printf("Successfully created %d keys for request: %s", numValidators, requestID)
	SetRequestStatus(requestID, "successful", keys)
}
