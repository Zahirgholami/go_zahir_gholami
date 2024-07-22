package main

import "fmt"

func main() {
	fmt.Println(getSequence(4))
	fmt.Println(getSequence(15))
	fmt.Println(getSequence(100))
}

func getSequence(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}
