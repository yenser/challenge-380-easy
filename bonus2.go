package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func bonus2() {
	file, err := os.Open("./enable1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	start := time.Now()

	found := false

	for scanner.Scan() {
		word := scanner.Text()

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