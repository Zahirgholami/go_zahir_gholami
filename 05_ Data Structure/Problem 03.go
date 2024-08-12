package main

import (
	"fmt"
	"sort"
)

func main() {
	data1 := []float64{1, 1, 2, 5, 6, 8, 12, 4, 5, 5, 5, 8, 9}

	fmt.Printf("sum: %.2f\n", sum(data1))       // 71.00
	fmt.Printf("mean: %.2f\n", mean(data1))     // 5.46
	fmt.Printf("median: %.2f\n", median(data1)) // 5.00
	fmt.Printf("mode: %.2f\n", mode(data1))     // 5.00

	data2 := []float64{6, 7, 1, 11, 8, 12, 6, 12, 1, 7, 8, 10, 14}

	fmt.Printf("sum: %.2f\n", sum(data2))       // 103.00
	fmt.Printf("mean: %.2f\n", mean(data2))     // 7.92
	fmt.Printf("median: %.2f\n", median(data2)) // 8.00
	fmt.Printf("mode: %.2f\n", mode(data2))     // 1.00
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
	if n%2 == 1 {
		return data[n/2]
	}
	return (data[n/2-1] + data[n/2]) / 2
}

func mode(data []float64) float64 {
	frequency := make(map[float64]int)
	for _, num := range data {
		frequency[num]++
	}

	modeVal := data[0]
	maxCount := frequency[modeVal]

	for num, count := range frequency {
		if count > maxCount {
			maxCount = count
			modeVal = num
		}
	}
	return modeVal
}
