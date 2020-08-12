package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func getResponseSize(url string) {

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
	startTime := time.Now()
	getResponseSize("http://www.google.com")
	getResponseSize("http://www.instagram.com")
	getResponseSize("http://www.stackoverflow.com")

	endTime := time.Now()
	fmt.Println("Elapsed time :", endTime.Sub(startTime))
}
