package main

import "fmt"

type Customer struct {
	name string
	age  int
}

func main() {
	var c Customer

	cPtr := &c
	fmt.Println(cPtr)

	var c2Ptr *Customer
	fmt.Println(c2Ptr)

	c2Ptr = new(Customer)
	fmt.Println(*c2Ptr)

}
