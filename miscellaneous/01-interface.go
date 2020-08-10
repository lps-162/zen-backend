package main

import "fmt"

type Order struct {
	OrderNo  int
	OrderAmt int
}

type Customer struct {
	CustomerNumber int
	CustomerName   string
	Phone          string
	CreditLimit    int
}

func whatType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("Int type")
	case Customer:
		fmt.Println("Customer")
	case Order:
		fmt.Println("Order")
	case []Order:
		fmt.Println("Slice Pointer")
	default:
		fmt.Println("Type Not Found")
	}
}

func main() {
	var i int
	var c Customer
	var o Order
	var s []Order

	var emp interface{}
	var myFunc func(int) int
	var ptr *int
	fmt.Println(emp)
	fmt.Println(myFunc)
	fmt.Println(ptr)
	emp = i
	emp = c
	emp = o
	emp = s
	fmt.Println(emp)

	fmt.Println("What types:")
	fmt.Println("-----------")
	whatType(i)
	whatType(c)
	whatType(o)
	whatType(s)
	fmt.Println()
}
