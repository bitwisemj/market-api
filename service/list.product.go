package service

import (
	"log"

	"github.com/bitwisemj/api-market/config"
	"github.com/bitwisemj/api-market/entities"
)

func GetProducts(page int, size int) ([]entities.Product, error) {

	log.Printf("Start getting products on page %v and size %v", page, size)
	db, err := config.GetConnection()

	if err != nil {
		log.Fatalf("Could not establish database connection %v", err)
		return nil, err
	}

	products := []entities.Product{}
	db.Offset(page).Limit(size).Find(&products)

	log.Print("Got products successfully")
	return products, nil
}
