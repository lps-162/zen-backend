package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Customer struct {
	CustomerNumber int
	CustomerName   string
	Phone          string
	CreditLimit    int
}

var customers = []Customer{
	Customer{CustomerNumber: 101, CustomerName: "John Richard",
		Phone: "999-11-888", CreditLimit: 500},
	Customer{CustomerNumber: 105, CustomerName: "Shawn Mendes",
		Phone: "999-11-777", CreditLimit: 200},
	Customer{CustomerNumber: 103, CustomerName: "Thomas Edison",
		Phone: "999-11-555", CreditLimit: 300},
}

func main() {

	file, err := json.MarshalIndent(customers, "", "\t")
	if err != nil {
		log.Fatal("Couldnt convert to json format")
	}

	fileName := "customers.json"
	err = ioutil.WriteFile(fileName, file, 0777)
	if err != nil {
		log.Fatal("Error writing to the file")
	}
}
