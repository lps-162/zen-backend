package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"zentools-mysql/models"
	"zentools-mysql/services"

	"github.com/gorilla/mux"
)

func (h *Handler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	service := services.Service{}
	customers, err := service.GetCustomers(h.DBConn)
	if err != nil {
		fmt.Fprintf(w, "Couldnt fetch customers")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h *Handler) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	strVars := mux.Vars(r)
	customerNumber, err := strconv.Atoi(strVars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Customer Number Invalid")
		return
	}
	// validations are done

	service := services.Service{}
	customer, err := service.GetCustomerByNumber(h.DBConn, customerNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Customer Not Found : %d", customerNumber)
			return
		}

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
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

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Customer created successfully")
}

func (h *Handler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Updating customer")
}

func (h *Handler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Deleting customer")
}
