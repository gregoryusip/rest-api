package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"rest-api/models"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id, omitempty"`
	Message string `json:"message, omitempty"`
}

type Response struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    []models.Product `json:"data"`
}

func AddProduct(w http.ResponseWriter, r *http.Request) {

	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		log.Fatalf("Can't decode from the request body. %v", err)
	}

	insertID := models.AddProduct(product)

	res := response{
		ID:      insertID,
		Message: "Book's Data already added",
	}

	json.NewEncoder(w).Encode(res)

}

func TakeProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Can't change from string to int. %v", err)
	}

	product, err := models.TakeProduct(int64(id))

	if err != nil {
		log.Fatalf("Can't take the book's data. %v", err)
	}

	json.NewEncoder(w).Encode(product)

}

func TakeAllProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	products, err := models.TakeAllProducts()

	if err != nil {
		log.Fatalf("Can't take the data. %v", err)
	}

	var response Response
	response.Status = 1
	response.Message = "Success"
	response.Data = products

	json.NewEncoder(w).Encode(response)

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Can't change to string to int. %v", err)
	}

	var product models.Product

	err = json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		log.Fatalf("Can't decode request body. %v", err)
	}

	updatedRows := models.UpdateProduct(int64(id), product)

	msg := fmt.Sprintf("Success to update the Book. The amount that was updated was %v rows/record", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Can't change from string to int. %v", err)
	}

	deletedRows := models.DeleteProduct(int64(id))

	msg := fmt.Sprintf("Success to delete the book. The amount that was deleted was %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)

}
