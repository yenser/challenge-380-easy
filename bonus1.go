package main

import (
	"fmt"
	"time"
)

func bonus1(words []string) {

	var m = make(map[string][]string)

	start := time.Now()

	for _, word := range words {
		code := getMorse(word)

		_, ok := m[code]

		if !ok {
			m[code] = []string{word}
		} else {
			m[code] = append(m[code], word)
		}
	}

	for key, value := range m {
		if len(value) == 13 {
			fmt.Println("Key:", key, "Value:", value)
			break
		}
	}

	timeTrack(start, "Bonus 1")
}