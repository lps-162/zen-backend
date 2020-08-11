package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"zentools-mysql/handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	database = "zentools"
	user     = "root"
	password = "sathyam"
)

func setupDatabaseConnection() *sql.DB {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, database)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {

	db := setupDatabaseConnection()

	router := mux.NewRouter()

	h := handlers.Handler{DBConn: db}

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
