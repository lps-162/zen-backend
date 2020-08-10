package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of all customers")
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Creating a new customer")
}
func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Customer Details of Id ")
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Updating customer")
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Deleting customer")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/customers", GetCustomers).Methods("GET")
	router.HandleFunc("/customers", CreateCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", GetCustomerById).Methods("GET")
	router.HandleFunc("/customers/{id}", UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", DeleteCustomer).Methods("DELETE")

	fmt.Println("Server listening on port no 8080")
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal("Error starting up server", err)
	}
}
