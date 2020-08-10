package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var customers = []Customer{
	Customer{CustomerNumber: 101, CustomerName: "John Richard",
		Phone: "999-11-888", CreditLimit: 500},
	Customer{CustomerNumber: 105, CustomerName: "Shawn Mendes",
		Phone: "999-11-777", CreditLimit: 200},
	Customer{CustomerNumber: 103, CustomerName: "Thomas Edison",
		Phone: "999-11-555", CreditLimit: 300},
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	// log.Println("Customers List :", customers)
	// fmt.Fprintf(w, "List of all customers")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)

}

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	strVars := mux.Vars(r)
	customerNumber, err := strconv.Atoi(strVars["id"])

	if err != nil {
		fmt.Fprintf(w, "Customer Number Invalid")
		return
	}

	for _, c := range customers {
		if c.CustomerNumber == customerNumber {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(c)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Customer Not Found")
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Creating a new customer")
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Updating customer")
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Deleting customer")
}
