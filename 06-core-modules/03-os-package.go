package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	numbers := os.Args[1:]
	fmt.Println(numbers)

	fileName := "sample.txt"
	custoFile, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer custoFile.Close()

	fmt.Println(os.Getenv("GOPATH"))

	for _, e := range os.Environ() {
		fmt.Println(e)
	}
}
