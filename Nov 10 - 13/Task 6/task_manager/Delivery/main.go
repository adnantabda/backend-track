package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"task_manager/Delivery/controllers"
	"task_manager/Delivery/routers"
	"task_manager/Domain"
	"task_manager/Infrastructure"
	repositories "task_manager/Repositories"
	"task_manager/Usecases"
)

func main() {
	// Initialize MongoDB connection
	mongoURI := getEnv("MONGODB_URI", "mongodb://localhost:27017")
	dbName := getEnv("DB_NAME", "taskdb")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the MongoDB server to verify the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	db := client.Database(dbName)

	// Initialize repositories
	taskRepo := repositories.NewTaskRepositoryMongo(db)
	userRepo := repositories.NewUserRepositoryMongo(db)

	// Initialize services
	passwordSvc := infrastructure.NewPasswordService()
	jwtSecret := getEnv("JWT_SECRET", "your-secret-key")
	jwtService := infrastructure.NewJWTService(jwtSecret)

	// Initialize use cases
	taskUseCase := Usecases.NewTaskUseCase(taskRepo)
	userUseCase := Usecases.NewUserUseCase(userRepo, passwordSvc, jwtService)

	// Initialize controllers
	taskController := controllers.NewTaskController(taskUseCase)
	userController := controllers.NewUserController(userUseCase)

	// Setup router
	r := routers.SetupRouter(taskController, userController, jwtService)

	// Start server in a goroutine
	port := getEnv("PORT", "8080")
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		log.Printf("Server is running on http://localhost:%s\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	// Disconnect from MongoDB
	if err := client.Disconnect(context.Background()); err != nil {
		log.Fatalf("Failed to disconnect from MongoDB: %v", err)
	}

	log.Println("Server exiting")
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
