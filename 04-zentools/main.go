package main

import (
	"fmt"
	"log"
	"net/http"
	"zentools/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	h := handlers.Handler{}

	router.HandleFunc("/customers", h.GetCustomers).Methods("GET")
	router.HandleFunc("/customers", h.CreateCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", h.GetCustomerById).Methods("GET")
	router.HandleFunc("/customers/{id}", h.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", h.DeleteCustomer).Methods("DELETE")

	router.HandleFunc("/orders", h.GetOrders).Methods("GET")

	fmt.Println("Server listening on port no 8080")
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal("Error starting up server", err)
	}
}
