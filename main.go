package main

import (
	"fmt"
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

	bonus2()
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