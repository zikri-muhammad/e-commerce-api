package services

import (
	"e-commerce-api/internal/repositories"
)

// ProductService handles business logic for products
type ProductService struct {
	productRepo repositories.ProductRepository
}

// NewProductService creates a new ProductService
func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return ProductService{productRepo: productRepo}
}

// GetAllProducts gets all products using the repository
func (s *ProductService) GetAllProducts() ([]map[string]interface{}, error) {
	// Add any business logic here if needed (e.g., filtering, transformations)
	return s.productRepo.GetAll()
}
