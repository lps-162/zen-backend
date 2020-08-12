package main

import "fmt"

func main() {

	var action int

	fmt.Scanln(&action)

	defer func() {
		panickedOn := recover()
		if panickedOn != nil {
			fmt.Println("Recovering :", panickedOn)
		}
	}()

	switch action {
	case 1:
		fmt.Println("Listing all customers")
	case 2:
		fmt.Println("List of all products")
	default:
		panic("Improper Action")
	}
}
