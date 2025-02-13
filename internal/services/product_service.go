package services

import (
	"e-commerce-api/internal/models"
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

// CreateProduct creates a new product
func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.productRepo.Create(product)
}

// GetProductByID gets a product by its ID
func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	return s.productRepo.GetByID(id)
}

// UpdateProduct updates an existing product
func (s *ProductService) UpdateProduct(product *models.Product) error {
	return s.productRepo.Update(product)
}

// DeleteProduct deletes a product by its ID
func (s *ProductService) DeleteProduct(id uint) error {
	return s.productRepo.Delete(id)
}
