package service

import (
	"log"

	"github.com/bitwisemj/api-market/config"
	"github.com/bitwisemj/api-market/dto"
	"github.com/bitwisemj/api-market/entities"
)

func UpdateProduct(productId int64, productDTO *dto.ProductDTO) error {

	log.Printf("Update product with id %v and data %v", productId, productDTO)

	db, err := config.GetConnection()

	if err != nil {
		log.Fatalf("Could not establish database connection")
		return err
	}

	product := entities.Product{}
	db.First(&product, productId)

	product.Name = productDTO.Name
	product.Price = productDTO.Price
	product.Quantity = productDTO.Quantity
	db.Save(&product)

	log.Printf("Product with id %v updated successfully", productId)
	return nil
}
