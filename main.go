package main

import (
	"fmt"
	"strings"
)
var morseCode = [26]string {".-",    "-...",  "-.-.",  "-..",
								".",     "..-.",  "--.",   "....",
								"..",    ".---",  "-.-",   ".-..",
								"--",    "-.",    "---",   ".--.",
								"--.-",  ".-.",   "...",   "-",
								"..-",   "...-",  ".--",   "-..-",
								"-.--",  "--.."}

func main() {
	smorse("sos")
	smorse("daily")
	smorse("programmer")
	smorse("bits")
	smorse("three")
}

func smorse(word string) {
	
	lowerstringWord := strings.ToLower(word);

	code := ""

	for i := 0; i < len(lowerstringWord); i++ {
		code += morseCode[lowerstringWord[i]-97]
	}
	
	fmt.Printf("%s => %s\n", word, code);
}