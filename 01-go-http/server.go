package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		fmt.Fprintf(w, "Method is not allowed")
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, "URL not handled", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello from Golang")
}

func oneMoreHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "One more handler provided")
}

func main() {
	fmt.Println("Starting Server on Port no 8080")

	http.HandleFunc("/onemore", oneMoreHandler)
	http.HandleFunc("/", helloHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Server failed, Exiting")
		os.Exit(-1)
	}

}
