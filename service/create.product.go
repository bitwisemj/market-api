package service

import (
	"log"

	"github.com/bitwisemj/api-market/config"
	"github.com/bitwisemj/api-market/dto"
	"github.com/bitwisemj/api-market/entities"
)

func CreateProduct(productDTO *dto.ProductDTO) error {

	log.Printf("A new product will be created with data %v", productDTO)

	product := entities.Product{
		Name:     productDTO.Name,
		Price:    productDTO.Price,
		Quantity: productDTO.Quantity,
	}

	db, err := config.GetConnection()

	if err != nil {
		log.Fatalf("Could not establish database connection: %v", err)
		return err
	}

	db.Create(&product)
	log.Printf("Product created successfully with id %v", product.ID)
	return nil
}
