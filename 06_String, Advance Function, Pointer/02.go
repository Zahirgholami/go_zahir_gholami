package main

import (
	"fmt"
)

func main() {
	fmt.Println(spinString("alta"))    // aalt
	fmt.Println(spinString("alterra")) // aalrtre
	fmt.Println(spinString("sepulsa")) // saesplu
}

func spinString(word string) string {
	if len(word) < 2 {
		return word
	}
	rotatedWord := string(word[0]) + word[2:] + string(word[1])
	return rotatedWord
}
