package main

import "fmt"

func calcTotal(a int, b int, tch chan int, ch2 chan int) {
	sum := a + b
	tch <- sum
	ch2 <- sum
}

func main() {

	tch := make(chan int)
	ch2 := make(chan int)

	go calcTotal(19, 24, tch, ch2)

	result := <-tch
	res2 := <-ch2
	fmt.Println("Total is :", result, res2)
}
