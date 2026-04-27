package main

import (
	config "kilosale_main/config"
)

func main() {
	// Load environment variables
	config.LoadConfig()

	// // Initialize database connection
	// db := database.DBInit()
	// defer db.Close()

	// // Initialize repositories
	// userRepo := userRepo.NewUserRepository(db)

	// // Initialize services
	// userService := userService.NewUserService(userRepo)

	// // Initialize handlers
	// userHandler := userHandler.NewUserHandler(userService)
}
