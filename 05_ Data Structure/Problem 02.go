package main

import "fmt"

func main() {
	fmt.Println(primeSum([]int{1, 2, 4, 5, 12, 19, 30})) // 26
	fmt.Println(primeSum([]int{45, 17, 44, 67, 11}))     // 95
	fmt.Println(primeSum([]int{15, 12, 9, 10}))          // 0
}

func primeSum(numbers []int) int {
	isPrime := func(n int) bool {
		if n <= 1 {
			return false
		}
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	sum := 0
	for _, number := range numbers {
		if isPrime(number) {
			sum += number
		}
	}

	return sum
}
