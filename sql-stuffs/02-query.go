package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:sathyam@/classicmodels")
	defer db.Close()

	if err != nil {
		log.Fatal("Error occured while setting up mysql connection", err)
	}

	sqlQuery := "select count(*) from customers"

	var oneMoreCount int
	err = db.QueryRow(sqlQuery).Scan(&oneMoreCount)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Using QueryRow :", oneMoreCount)

	// result, err := db.Query(sqlQuery)
	// defer result.Close()

	// if err != nil {
	// 	panic(err.Error())
	// }

	// for result.Next() {
	// 	var rowCount string
	// 	err := result.Scan(&rowCount)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	fmt.Println("Total number of rows in customers table :", rowCount)
	// }
}
