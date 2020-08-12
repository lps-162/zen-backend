package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func getResponseSize(url string) {
	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done :", url, len(body))
}

func main() {
	wg.Add(3)
	startTime := time.Now()
	go getResponseSize("http://www.google.com")
	go getResponseSize("http://www.instagram.com")
	go getResponseSize("http://www.stackoverflow.com")

	wg.Wait()
	endTime := time.Now()
	fmt.Println("Elapsed time :", endTime.Sub(start))
}
