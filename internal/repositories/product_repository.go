package repositories

// ProductRepository is responsible for interacting with the database or data source
type ProductRepository struct {
	// You can add DB connection here in a real implementation
}

// NewProductRepository creates a new ProductRepository
func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

// GetAll returns a list of all products (mocked data)
func (r *ProductRepository) GetAll() ([]map[string]interface{}, error) {
	// In a real-world scenario, this would be a database query.
	// Here we're returning a mocked list of products.
	products := []map[string]interface{}{
		{
			"id":          1,
			"name":        "Kopi Kapal Api",
			"description": "kopi kapal api adalah sebuah kopi hitam",
			"price":       9000,
			"stock":       9,
			"categor_id":  1,
			"created_at":  "14:32 10/01/2025",
			"updated_at":  "14:32 10/01/2025",
		},
		{
			"id":          2,
			"name":        "Kopi Luak",
			"description": "kopi kapal api adalah sebuah kopi hitam",
			"price":       9000,
			"stock":       9,
			"categor_id":  1,
			"created_at":  "14:32 10/01/2025",
			"updated_at":  "14:32 10/01/2025",
		},
	}
	return products, nil
}
