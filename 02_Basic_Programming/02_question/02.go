package main

import (
	"fmt"
)

func getGrade(score float64) string {
	if score < 0 || score > 100 {
		return "Invalid score"
	} else if score >= 85 {
		return "A"
	} else if score >= 70 {
		return "B"
	} else if score >= 55 {
		return "C"
	} else if score >= 40 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	var score float64
	fmt.Print("Enter the score:80 ")
	fmt.Scan(&score)
	grade := getGrade(score)
	fmt.Println("The grade for a score of 85 is A", grade)
}
