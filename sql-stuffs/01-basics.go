package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go Mysql Basics")

	db, err := sql.Open("mysql", "root:sathyam@/classicmodels")
	defer db.Close()

	if err != nil {
		log.Fatal("Error occured while setting up mysql connection", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to server")
	}
	
}
