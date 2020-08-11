package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	db, err := sql.Open("mysql", "root:sathyam@/classicmodels")
	defer db.Close()

	if err != nil {
		log.Fatal("Error occured while setting up mysql connection", err)
	}

	sqlQuery := "select customerNumber, customerName, phone, creditLimit from customers limit 10"

	result, err := db.Query(sqlQuery)
	defer result.Close()

	checkError(err)

	for result.Next() {
		var customerNumber int
		var customerName string
		var phone string
		var creditLimit float64

		err := result.Scan(&customerNumber, &customerName, &phone, &creditLimit)

		checkError(err)
		fmt.Println(customerNumber, customerName, phone, creditLimit)

	}
}

func sum(a int, b int) (int, error) {

	if a == 0 || b == 0 {
		err := errors.New("Invalid a and b")
		return -1, err
	}

	sum := a + b
	return sum, nil
}
