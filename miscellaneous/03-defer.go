package main

import "fmt"

func first() {
	fmt.Println("First")
}

func second() {
	fmt.Println("Second")
}

func cleanup() {
	fmt.Println("Cleanup")
}

func main() {
	defer cleanup()
	defer first()
	second()
}
