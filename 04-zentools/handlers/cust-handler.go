package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"zentools/models"

	"github.com/gorilla/mux"
)

var customers = []models.Customer{
	models.Customer{CustomerNumber: 101, CustomerName: "John Richard",
		Phone: "999-11-888", CreditLimit: 500},
	models.Customer{CustomerNumber: 105, CustomerName: "Shawn Mendes",
		Phone: "999-11-777", CreditLimit: 200},
	models.Customer{CustomerNumber: 103, CustomerName: "Thomas Edison",
		Phone: "999-11-555", CreditLimit: 300},
}

func (h *Handler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	// r.URL.Query()
	// log.Println("Customers List :", customers)
	// fmt.Fprintf(w, "List of all customers")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)

}

func (h *Handler) GetCustomerById(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var customer models.Customer
	err := json.Unmarshal(reqBody, &customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Badly formatted data")
		return
	}

	customers = append(customers, customer)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Customer created successfully")
}

func (h *Handler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Updating customer")
}

func (h *Handler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Deleting customer")
}
