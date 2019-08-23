package main

import (
	"fmt"
	"time"
)

func bonus4(words []string) {
	
	start := time.Now()

	for _, word := range words {
		if len(word) == 13 && isPalindrome(getMorse(word)) {
			fmt.Println("13 letter word palindrome", word)
			break
		}
	}


	timeTrack(start, "Bonus 4")
}