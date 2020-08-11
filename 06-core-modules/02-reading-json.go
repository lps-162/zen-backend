package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Customer struct {
	CustomerNumber int
	CustomerName   string
	Phone          string
	CreditLimit    int
}

func main() {

	file, err := ioutil.ReadFile("customers.json")

	if err != nil {
		log.Fatal("Error reading json file")
	}

	customers := []Customer{}

	err = json.Unmarshal(file, &customers)
	if err != nil {
		log.Fatal("Error decoding/unmarshaling json file")
	}

	fmt.Println("Customer Details")
	fmt.Println("----------------")
	for _, c := range customers {
		fmt.Println(c)
	}
}
