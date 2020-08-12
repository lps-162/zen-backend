package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sum(a int, b int) {
	defer wg.Done()
	fmt.Println("Result is :", a+b)
}

func main() {
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go sum(i, 0)
	}

	fmt.Println("End of Main")

	wg.Wait()
}
