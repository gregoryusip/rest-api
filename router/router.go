package router

import (
	"rest-api/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/product", controller.TakeAllProducts).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/product/{id}", controller.TakeProduct).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/product", controller.AddProduct).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/product/{id}", controller.UpdateProduct).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/product", controller.DeleteProduct).Methods("DELETE", "OPTIONS")

	return router
}
