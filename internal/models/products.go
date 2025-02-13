package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Stok        float32 `json:"stok"`
	Description string  `json:"description"`
}
