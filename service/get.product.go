package service

import (
	"log"

	"github.com/bitwisemj/api-market/config"
	"github.com/bitwisemj/api-market/entities"
)

func GetProduct(productId int64) (entities.Product, error) {

	db, err := config.GetConnection()
	product := entities.Product{}

	if err != nil {
		log.Fatalf("Could not establish database connection: %v", err)
		return product, err
	}

	db.Where("id = ?", productId).Find(&product)
	log.Printf("Got product with id %v successfully", productId)
	return product, nil
}
