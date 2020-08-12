package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	go generateNums(10, ch)

	go worker(ch)
	go worker(ch)
	go worker(ch)

	for i := range ch {
		fmt.Println("Main worker received value :", i)
		time.Sleep(1 * time.Second)
	}

}
func worker(ch chan int) {
	for i := range ch {
		fmt.Println("Worker ", i)
		time.Sleep(1 * time.Second)
	}
}

func generateNums(q int, ch chan int) {
	for i := 0; i < q; i++ {
		ch <- i
	}
	close(ch)
}
