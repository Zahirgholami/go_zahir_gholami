package main

import (
	"fmt"
	"sort"
)

func main() {
	data1 := []float64{1, 1, 2, 5, 6, 8, 12, 4, 5, 5, 5, 8, 9}
	printStatistics(data1) // sum: 71.00, mean: 5.46, median: 5.00, mode: 5.00

	data2 := []float64{6, 7, 1, 11, 8, 12, 6, 12, 1, 7, 8, 10, 14}
	printStatistics(data2) // sum: 103.00, mean: 7.92, median: 8.00, mode: 1.00
}

func printStatistics(data []float64) {
	fmt.Printf("sum: %.2f\n", sum(data))
	fmt.Printf("mean: %.2f\n", mean(data))
	fmt.Printf("median: %.2f\n", median(data))
	fmt.Printf("mode: %.2f\n", mode(data))
}

func sum(data []float64) float64 {
	total := 0.0
	for _, num := range data {
		total += num
	}
	return total
}

func mean(data []float64) float64 {
	return sum(data) / float64(len(data))
}

func median(data []float64) float64 {
	sort.Float64s(data)
	n := len(data)
	if n%2 == 0 {
		return (data[n/2-1] + data[n/2]) / 2.0
	}
	return data[n/2]
}

func mode(data []float64) float64 {
	counts := make(map[float64]int)
	for _, num := range data {
		counts[num]++
	}

	maxCount := 0
	mode := data[0]
	for num, count := range counts {
		if count > maxCount {
			maxCount = count
			mode = num
		}
	}

	return mode
}
