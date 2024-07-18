package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(pickVocals("alterra academy"))    
  fmt.Println(pickVocals("i love programming"))
  fmt.Println(pickVocals("go is awesome programming language"))
}

func pickVocals(sentence string) string {
  vowels := "aeiouAEIOU"
  words := strings.Fields(sentence)
  var result []string
  for _, word := range words {
    var vocalWord string
    for _, char := range word {
      if strings.ContainsRune(vowels, char) {
        vocalWord += string(char)
      }
    }
    result = append(result, vocalWord)
  }
  return strings.Join(result, " ")
}

