package repositories

import (
	"e-commerce-api/internal/models"

	"gorm.io/gorm"
)

// ProductRepository is responsible for interacting with the database
type ProductRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new ProductRepository
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

// GetAll returns a list of all products from database
func (r *ProductRepository) GetAll() ([]map[string]interface{}, error) {
	var products []models.Product

	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}

	// Convert to map format for consistency with existing response
	result := make([]map[string]interface{}, len(products))
	for i, product := range products {
		result[i] = map[string]interface{}{
			"id":          product.ID,
			"name":        product.Name,
			"price":       product.Price,
			"stok":        product.Stok,
			"description": product.Description,
		}
	}

	return result, nil
}

// Create adds a new product to database
func (r *ProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

// GetByID retrieves a product by its ID
func (r *ProductRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// Update modifies an existing product
func (r *ProductRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

// Delete removes a product from database
func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}
