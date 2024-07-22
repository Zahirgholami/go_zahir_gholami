package main

import "fmt"

func main() {
	fmt.Println(spinString("alta"))    // aalt
	fmt.Println(spinString("alterra")) // aalrtre
	fmt.Println(spinString("sepulsa")) // saesplu
}

func spinString(sentence string) string {
	chars := []rune(sentence)
	if len(chars) < 2 {
		return sentence
	}

	result := []rune{}
	half := len(chars) / 2

	for i := 0; i < half; i++ {
		result = append(result, chars[i])
		result = append(result, chars[len(chars)-1-i])
	}

	if len(chars)%2 != 0 {
		result = append(result, chars[half])
	}

	return string(result)
}
