package main

import (
	"fmt"
	"strings"
)

func main() {

	primes := [10]int{2, 3, 5, 7, 11, 13, 17, 19}

	primeSlice := primes[3:6]

	fmt.Println(primes)
	fmt.Println(primeSlice)

	cities := [3]string{"goa", "cbe", "kochi"}

	cities2 := cities
	cities3 := &cities

	cities[0] = "Trivandrum"

	fmt.Println(cities)
	fmt.Println(cities2)
	fmt.Println(cities3)

	sentence := strings.Join(cities[:], ":")
	fmt.Println(sentence)
}
