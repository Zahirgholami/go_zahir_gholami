package main

import (
	"fmt"
)

func main() {
	// Test Case 1
	input1 := []int{1, 2, 4, 5, 12, 19, 30}
	fmt.Println(primeSum(input1)) // 26

	// Test Case 2
	input2 := []int{45, 17, 44, 67, 11}
	fmt.Println(primeSum(input2)) // 95

	// Test Case 3
	input3 := []int{15, 12, 9, 10}
	fmt.Println(primeSum(input3)) // 0
}

func primeSum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		if isPrime(num) {
			sum += num
		}
	}
	return sum
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
