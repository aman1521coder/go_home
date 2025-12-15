package main

import (
	"log"
	"net/http"

	database "primeauction/api/Database"
	repository "primeauction/api/Repository"
	"primeauction/api/handler"
	"primeauction/api/routes"
	"primeauction/api/service"
)

func main() {
	// Initialize database
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.CloseDB()

	// Initialize repositories
	itemRepo := repository.NewItemRepository(database.DB)
	userRepo := repository.NewUserRepository(database.DB)

	// Initialize services
	itemService := service.NewItemService(itemRepo)
	userService := service.NewUserService(userRepo)

	// Initialize handlers
	itemHandler := handler.NewItemHandler(itemService)
	userHandler := handler.NewUserHandler(userService)

	// Setup and register routes
	routesList := routes.SetupRoutes(itemHandler, userHandler)
	routes.RegisterRoutes(&routesList)

	// Start server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
