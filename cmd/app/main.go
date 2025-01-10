package main

import (
	"e-commerce-api/internal/handlers" // Import the handlers package
	"e-commerce-api/internal/repositories"
	"e-commerce-api/internal/routers"
	"e-commerce-api/internal/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()
	// Setup the repository (could be a DB in the future)
	productRepo := repositories.NewProductRepository()

	// Setup the service, injecting the repository
	productService := services.NewProductService(*productRepo)

	// Setup the handler, injecting the service
	productHandler := handlers.NewProductHandler(productService)
	// Setup routes
	routers.SetupProductRoutes(app, productHandler)

	// Run the server
	log.Fatal(app.Listen(":3000"))
}
