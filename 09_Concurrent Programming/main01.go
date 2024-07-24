package main

import (
	"fmt"
	"sync"
	"time"
)

func reverseWord(word string) {
	var wg sync.WaitGroup
	wg.Add(len(word))

	for i := len(word) - 1; i >= 0; i-- {
		go func(j int) {
			defer wg.Done()
			time.Sleep(3 * time.Second)
			fmt.Println(word[j:])
		}(i)
	}

	wg.Wait()
}

func main() {
	word := "phone"
	reverseWord(word)
}
