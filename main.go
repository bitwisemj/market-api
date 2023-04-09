package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bitwisemj/api-market/api"
	"github.com/bitwisemj/api-market/config"
)

func main() {

	config.Start()
	os.Setenv("APP_PORT", ":8082")
	port := os.Getenv("APP_PORT")
	server := http.NewServeMux()
	server.HandleFunc("/products", api.GetProducts)
	server.HandleFunc("/products/create", api.CreateProduct)
	server.HandleFunc("/products/find", api.GetProduct)
	server.HandleFunc("/products/delete", api.DeleteProduct)
	server.HandleFunc("/products/update", api.UpdateProduct)

	log.Printf("Starting server at port %v", port)
	http.ListenAndServe(port, server)
}
