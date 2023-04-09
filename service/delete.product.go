package service

import (
	"log"

	"github.com/bitwisemj/api-market/config"
	"github.com/bitwisemj/api-market/entities"
)

func DeleteProduct(productId int64) error {

	log.Printf("Product with id %v will be deleted", productId)
	db, err := config.GetConnection()

	if err != nil {
		log.Fatal("Could not establish database connection")
		return err
	}

	product := entities.Product{}
	db.First(&product, productId)
	db.Delete(&product)

	log.Printf("Product with id %v deleted successfully", productId)

	return nil
}
