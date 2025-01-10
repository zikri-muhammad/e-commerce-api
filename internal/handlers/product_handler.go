package handlers

import (
	"e-commerce-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

// ProductHandler handles product-related requests
type ProductHandler struct {
	productService services.ProductService
}

// NewProductHandler creates a new ProductHandler
func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// GetProducts handles GET /products request
func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	// Call the service to get products
	products, err := h.productService.GetAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to fetch products"})
	}
	return c.JSON(products)
}
