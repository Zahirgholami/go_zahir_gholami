package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(pickVocals("alterra academy"))                    // aea aae
	fmt.Println(pickVocals("i love programming"))                 // i oe oai
	fmt.Println(pickVocals("go is awesome programming language")) // o i aeoe oai auae
}

func pickVocals(sentence string) string {
	vowels := "aeiou"
	words := strings.Fields(sentence)
	result := []string{}

	for _, word := range words {
		vocalLetters := ""
		for _, char := range word {
			if strings.ContainsRune(vowels, char) {
				vocalLetters += string(char)
			}
		}
		result = append(result, vocalLetters)
	}

	return strings.Join(result, " ")
}
