package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func InitRedis() {
	fmt.Println("Inside InitRedis() function...")
	redisHost := ""
	if isRunningInDocker() {
		redisHost = "redis:6379" // Docker's host mapping
	} else {
		redisHost = "localhost:6379" // local Redis
	}

	fmt.Println("Redis Host: " + redisHost)

	rdb = redis.NewClient(&redis.Options{
		Addr:     redisHost, // Redis server address
		Password: "",        // No password by default
		DB:       0,         // Default DB
	})

	// Test connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	log.Println("Connected to Redis!")
}

func SetRequestStatus(requestID, status string, keys []string) error {
	data, err := rdb.Get(ctx, requestID).Result()
	if err != nil {
		return err
	}

	var validatorRequest ValidatorRequest
	if err := json.Unmarshal([]byte(data), &validatorRequest); err != nil {
		return err 
	}

	// Update status & keys
	validatorRequest.Status = status
	if status == "successful" {
		validatorRequest.Keys = keys
	}

	// Save updated struct back to Redis
	updatedData, err := json.Marshal(validatorRequest)
	if err != nil {
		return err
	}

	return rdb.Set(ctx, requestID, updatedData, 24*time.Hour).Err()
}

func GetRequestStatus(requestID string) (*ValidatorRequest, error) {
	data, err := rdb.Get(ctx, requestID).Result()
	if err != nil {
		return nil, err
	}

	var validatorRequest ValidatorRequest
	if err := json.Unmarshal([]byte(data), &validatorRequest); err != nil {
		return nil, err
	}

	return &validatorRequest, nil
}

func isRunningInDocker() bool {
	_, err := os.Stat("/.dockerenv")
	return err == nil
}
