package main

import (
	"log"
	"task_manager/data"
	"task_manager/router"
)

func main() {
	// Initialize MongoDB connection
	err := data.InitMongo("mongodb://localhost:27017", "taskdb")
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer data.DisconnectMongo()

	r := router.SetupRouter()
	r.Run(":8080")
}
