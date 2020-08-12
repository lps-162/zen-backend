package main

import (
	"fmt"
	"time"
)

func sum(a int, b int) {
	fmt.Println("Result is :", a+b)
}

func main() {
	go sum(10, 20)
	go sum(11, 20)
	go sum(15, 60)

	fmt.Println("End of Main")

	time.Sleep(2 * time.Second)
}
