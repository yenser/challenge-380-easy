package main

import (
	"fmt"
	"strings"
	"time"
)

func getMorse(word string) string {
	lowerstringWord := strings.ToLower(word)

	code := ""

	for i := 0; i < len(lowerstringWord); i++ {
		code += morseCode[lowerstringWord[i]-97]
	}

	return code
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n\n", name, elapsed)
}
