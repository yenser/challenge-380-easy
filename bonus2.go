package main

import (
	"fmt"
	"time"
)

func bonus2(words []string) {
	start := time.Now()

	found := false

	for _, word := range words {
		if len(word) >= 4 {
			code := getMorse(word)

			inARow := 0

			for i := 0; i < len(code); i++ {
				if code[i] == '-' {
					inARow++

					if inARow == 15 {
						found = true
						fmt.Println(word)
						break
					}

				} else {
					inARow = 0

					if (len(code) - i) < 15 {
						continue
					}
				}
			}
		}

		if found {
			break
		}
	}

	timeTrack(start, "Bonus 2")
}