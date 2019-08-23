package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var morseCode = [26]string{".-", "-...", "-.-.", "-..",
	".", "..-.", "--.", "....",
	"..", ".---", "-.-", ".-..",
	"--", "-.", "---", ".--.",
	"--.-", ".-.", "...", "-",
	"..-", "...-", ".--", "-..-",
	"-.--", "--.."}

func main() {

	challenge()

	bonus1()
}

func challenge() {
	start := time.Now()

	smorse("sos")
	smorse("daily")
	smorse("programmer")
	smorse("bits")
	smorse("three")

	timeTrack(start, "Challenge")
}

func smorse(word string) {

	code := getMorse(word)

	fmt.Printf("%s => %s\n", word, code)
}

func bonus1() {

	var m = make(map[string][]string)

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

	for scanner.Scan() {
		word := scanner.Text()
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
