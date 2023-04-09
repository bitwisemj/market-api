package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/bitwisemj/api-market/dto"
	"github.com/bitwisemj/api-market/service"
)

func CreateProduct(response http.ResponseWriter, request *http.Request) {

	productDTO, err := getBody(request)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	service.CreateProduct(productDTO)
	response.WriteHeader(http.StatusCreated)
}

func UpdateProduct(response http.ResponseWriter, request *http.Request) {

	productId, err := strconv.ParseInt(request.URL.Query().Get("productId"), 10, 64)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
		return
	}

	productDTO, err2 := getBody(request)

	if err2 != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
		return
	}

	err3 := service.UpdateProduct(productId, productDTO)

	if err3 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	response.WriteHeader(http.StatusNoContent)
}

func DeleteProduct(response http.ResponseWriter, request *http.Request) {

	productId, err := strconv.ParseInt(request.URL.Query().Get("productId"), 10, 64)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	service.DeleteProduct(productId)
	response.WriteHeader(http.StatusOK)
}

func GetProducts(response http.ResponseWriter, request *http.Request) {

	page, err1 := strconv.Atoi(request.URL.Query().Get("page"))
	size, err2 := strconv.Atoi(request.URL.Query().Get("size"))

	if err1 != nil || err2 != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	products, err := service.GetProducts(page, size)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(response)
	encoder.Encode(products)
}

func GetProduct(response http.ResponseWriter, request *http.Request) {

	productId, err := strconv.ParseInt(request.URL.Query().Get("productId"), 10, 64)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err1 := service.GetProduct(productId)

	if err1 != nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	response.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(response)
	encoder.Encode(product)
}

func getBody(request *http.Request) (*dto.ProductDTO, error) {

	bytes, err := io.ReadAll(request.Body)

	if err != nil {
		log.Fatalf("Could not read request body: %v", err)
		return nil, err
	}

	productDTO := dto.ProductDTO{}
	err2 := json.Unmarshal(bytes, &productDTO)

	if err2 != nil {
		log.Fatalf("Could not parse json to object: %v", err2)
		return nil, err2
	}

	return &productDTO, nil
}
