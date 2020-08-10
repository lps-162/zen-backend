package main

type Customer struct {
	CustomerNumber int    `json:"customer_number,omitempty"`
	CustomerName   string `json:"customer_name,omitempty"`
	Phone          string `json:"phone,omitempty"`
	CreditLimit    int    `json:"credit_limit,omitempty"`
}
