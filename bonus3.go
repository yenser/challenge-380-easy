package main

import (
	"fmt"
	"time"
)

func bonus3(words []string) {
	
	start := time.Now()
	
	for _, word := range words {
		if len(word) == 21 && word != "counterdemonstrations" && isPerfectBalance(getMorse(word)) {
			fmt.Println("Other perfectly balanced word is", word)
			break;
		}
	}

	timeTrack(start, "Bonus 3")
}