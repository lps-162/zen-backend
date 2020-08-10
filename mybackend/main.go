package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	fmt.Println("Server on Router")
	fmt.Println(router)
}
