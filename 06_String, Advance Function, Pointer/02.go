package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(spinString("alta"))    // aalt
    fmt.Println(spinString("alterra")) // aalrtre
    fmt.Println(spinString("sepulsa"))
}

func spinString(word string) string {
    if len(word) < 2 {
        return word
    }
    rotatedWord := string(word[0]) + word[2:] + string(word[1])
    return rotatedWord
}
