package main

import (
	"fmt"
)

// Function to generate prime numbers from 1 to 100
func generatePrimes(primeChan chan int) {
	for num := 2; num <= 100; num++ {
		isPrime := true
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primeChan <- num
		}
	}
	close(primeChan)
}

// Function to calculate the square of each prime number
func calculateSquare(primeChan, squareChan chan int) {
	for prime := range primeChan {
		squareChan <- prime * prime
	}
	close(squareChan)
}

func main() {
	primeChan := make(chan int)
	squareChan := make(chan int)

	go generatePrimes(primeChan)
	go calculateSquare(primeChan, squareChan)

	for square := range squareChan {
		fmt.Println(square)
	}
}
