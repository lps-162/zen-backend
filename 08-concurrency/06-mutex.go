package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var counter int

func main() {
	wg.Add(3)

	go increment("Golang")
	go increment("React")
	go increment("GCloud")

	wg.Wait()
	mutex.Lock()
	fmt.Println("Counter : ", counter)
	mutex.Unlock()
}

func increment(str string) {
	wg.Done()

	mutex.Lock()
	{
		for i := 0; i < len(str); i++ {
			counter++
		}
	}
	mutex.Unlock()
}
