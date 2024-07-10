package main

import (
	"fmt"
	"math"
)

func primeNumber(number int) bool {
	if number <= 1 {
		return false
	}
	if number == 2 || number == 3 {
		return true
	}
	if number%2 == 0 || number%3 == 0 {
		return false
	}
	for i := 5; i <= int(math.Sqrt(float64(number))); i += 6 {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(1500450271)) // false
	fmt.Println(primeNumber(13))         // true
	fmt.Println(primeNumber(17))         // true
	fmt.Println(primeNumber(20))         // false
	fmt.Println(primeNumber(35))         // false
}
