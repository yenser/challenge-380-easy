package main

import (
	"fmt"
	"time"
	"log"
)

var morseCode = [26]string{".-", "-...", "-.-.", "-..",
							".", "..-.", "--.", "....",
							"..", ".---", "-.-", ".-..",
							"--", "-.", "---", ".--.",
							"--.-", ".-.", "...", "-",
							"..-", "...-", ".--", "-..-",
							"-.--", "--.."}

func main() {

	words, err := getListOfWordsFromFile("./enable1.txt")
	if err != nil {
		log.Fatal(err)
	}

	challenge()

	bonus1(words)

	bonus2(words)

	bonus3(words)

	bonus4(words)
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