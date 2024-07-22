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
	vowels := "aeiouAEIOU"
	words := strings.Fields(sentence)
	result := []string{}

	for _, word := range words {
		filteredWord := ""
		for _, char := range word {
			if strings.ContainsRune(vowels, char) {
				filteredWord += string(char)
			}
		}
		result = append(result, filteredWord)
	}

	return strings.Join(result, " ")
}
