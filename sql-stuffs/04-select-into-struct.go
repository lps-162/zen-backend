package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

type Customer struct {
	CustomerNumber int
	CustomerName   string
	Phone          string
	CreditLimit    float64
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
	var customers []Customer

	for result.Next() {
		var c Customer

		err := result.Scan(&c.CustomerNumber, &c.CustomerName, &c.Phone, &c.CreditLimit)

		checkError(err)

		customers = append(customers, c)
	}

	for _, c := range customers {
		fmt.Println(c)
	}
}
