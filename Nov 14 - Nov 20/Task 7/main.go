package main

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"task_manager/controllers"
	"task_manager/data"
	"task_manager/models"
	"task_manager/router"
)

func main() {
	// Set JWT secret key
	if os.Getenv("JWT_SECRET") == "" {
		os.Setenv("JWT_SECRET", "your-secret-key") // Change this in production
	}

	// Initialize database
	db, err := gorm.Open(sqlite.Open("task_manager.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate the schema
	if err := db.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize services
	userService := data.NewUserService(db)
	taskService := data.NewTaskService(db)

	// Initialize controllers
	authController := controllers.NewAuthController(userService)
	taskController := controllers.NewTaskController(taskService)

	// Create admin user if not exists
	createAdminIfNotExists(userService)

	// Initialize router
	r := router.SetupRouter(authController, taskController)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Server running on port %s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func createAdminIfNotExists(us *data.UserService) {
	// Check if admin exists
	_, err := us.GetUserByUsername("admin")
	if err == nil {
		return // Admin already exists
	}

	// Create admin user
	if _, err := us.CreateUser("admin", "admin123", models.AdminRole); err != nil {
		log.Println("Warning: Failed to create admin user:", err)
	} else {
		log.Println("Created default admin user (username: admin, password: admin123)")
	}
}
