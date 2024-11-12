package models

import (
	"gorm.io/gorm"

	db "github.com/nomannaq/webshop-api-golang/cmd/database"
)

type Product struct {
	gorm.Model
	ID          int    `gorm:"unique; not null" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	Price       int    `gorm:"not null" json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Quantity    int    `json:"quantity"`
}

func CreateProduct(product *Product) error {
	return db.DB.Create(product).Error
}
func GetProducts() ([]Product, error) {
	var products []Product
	if err := db.DB.Find((&products)).Error; err != nil {
		return nil, err
	}
	return products, nil

}
