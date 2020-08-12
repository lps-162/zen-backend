package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	CustomerNumber int
	CustomerName   string
	Phone          string
	CreditLimit    float64
}

var customers = []Customer{
	Customer{CustomerNumber: 100, CustomerName: "Shivan",
		Phone: "999-888-99", CreditLimit: 1000},
	Customer{CustomerNumber: 101, CustomerName: "Shivan LP",
		Phone: "999-777-99", CreditLimit: 600},
	Customer{CustomerNumber: 102, CustomerName: "Shivan ",
		Phone: "999-555-99", CreditLimit: 500},
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
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
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Error starting up server", err)
	}
}
