package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	//Router *mux.Router
}

func (app *App) Run() {
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
