package main

import (
	"fmt"
)

func main() {
	// Test Case 1
	input1 := [][]int{
		{1, 2, 5, 4, 3},
		{1, 2, 7, 8},
	}
	fmt.Println(merge(input1)) // [1, 2, 5, 4, 3, 7, 8]

	// Test Case 2
	input2 := [][]int{
		{1, 2, 5, 4, 5},
		{7, 9, 10, 5},
	}
	fmt.Println(merge(input2)) // [1, 2, 5, 4, 7, 9, 10]

	// Test Case 3
	input3 := [][]int{
		{1, 4, 5},
		{7},
	}
	fmt.Println(merge(input3)) // [1, 4, 5, 7]
}

func merge(data [][]int) []int {
	uniqueElements := make(map[int]bool)
	result := []int{}

	for _, row := range data {
		for _, elem := range row {
			if !uniqueElements[elem] {
				uniqueElements[elem] = true
				result = append(result, elem)
			}
		}
	}

	return result
}
