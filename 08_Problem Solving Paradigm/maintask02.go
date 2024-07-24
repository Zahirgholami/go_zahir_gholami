package main

import "fmt"

func Frog(jumps []int) int {
	n := len(jumps)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = 0 // Starting at the first rock, cost is 0

	for i := 1; i < n; i++ {
		cost1 := dp[i-1] + abs(jumps[i]-jumps[i-1])
		cost2 := int(^uint(0) >> 1) // infinity

		if i > 1 {
			cost2 = dp[i-2] + abs(jumps[i]-jumps[i-2])
		}

		dp[i] = min(cost1, cost2)
	}

	return dp[n-1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
