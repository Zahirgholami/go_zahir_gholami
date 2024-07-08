package main

import (
	"fmt"
	"math"
)

func calculateTubeVolume(radius float64, height float64) float64 {
	phi := math.Pi
	volume := phi * math.Pow(radius, 2) * height
	return volume
}

func main() {
	radius := 5.0
	height := 10.0
	volume := calculateTubeVolume(radius, height)
	fmt.Println("The volume of the tube is 785.39 cubic units", volume)
}
