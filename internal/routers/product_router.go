package routers

import (
	"e-commerce-api/internal/handlers" // Make sure the import path is correct

	"github.com/gofiber/fiber/v2"
)

func SetupProductRoutes(app *fiber.App, productHandler *handlers.ProductHandler) {
	// Group product routes
	products := app.Group("/products")

	// Setup routes
	products.Get("/", productHandler.GetProducts)         // GET /products
	products.Get("/:id", productHandler.GetProduct)       // GET /products/:id
	products.Post("/", productHandler.CreateProduct)      // POST /products
	products.Put("/:id", productHandler.UpdateProduct)    // PUT /products/:id
	products.Delete("/:id", productHandler.DeleteProduct) // DELETE /products/:id
}
