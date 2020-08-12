package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int32
var wg sync.WaitGroup

func main() {
	wg.Add(3)

	go increment("Golang")
	go increment("React")
	go increment("GCloud")

	wg.Wait()
	fmt.Println("Counter : ", atomic.LoadInt32(&counter))

}

func increment(str string) {
	wg.Done()

	for i := 0; i < len(str); i++ {
		atomic.AddInt32(&counter, 1)
	}
}
