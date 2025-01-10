package routers

import (
	"e-commerce-api/internal/handlers" // Make sure the import path is correct

	"github.com/gofiber/fiber/v2"
)

func SetupProductRoutes(app *fiber.App, productHandler *handlers.ProductHandler) {
	app.Get("/products", productHandler.GetProducts)
	// app.Get("/product/:id", productHandler.GetProduct) // You can add more routes if needed
}
