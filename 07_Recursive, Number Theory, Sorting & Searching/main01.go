package main

import "fmt"

func main() {
	fmt.Println(fibX(5))
	fmt.Println(fibX(10))
}

func fibX(n int) int {
	if n <= 0 {
		return 0
	}

	fib := []int{0, 1}
	sum := 1

	for i := 2; i <= n; i++ {
		nextFib := fib[i-1] + fib[i-2]
		fib = append(fib, nextFib)
		sum += nextFib
	}

	return sum
}
