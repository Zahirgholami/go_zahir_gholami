package main

import (
	"fmt"
	"math"
)

func isComposite(n int) bool {
	if n < 4 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return true
		}
	}
	return false
}

func main() {
	// Channel for composite numbers
	composites := make(chan int)
	// Channel for squared values
	squares := make(chan int)

	go func() {
		for i := 1; i <= 100; i++ {
			if isComposite(i) {
				composites <- i
			}
		}
		close(composites)
	}()
	go func() {
		for num := range composites {
			squares <- num * num
		}
		close(squares)
	}()

	for square := range squares {
		if square%2 == 0 {
			fmt.Println("Pong")
		} else {
			fmt.Println("Ping")
		}
	}
}
