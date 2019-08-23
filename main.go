package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"log"
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

	bonus1()
}


func smorse(word string) {
	
	code := getMorse(word)
	
	fmt.Printf("%s => %s\n", word, code);
}

func getMorse(word string) string {
	lowerstringWord := strings.ToLower(word);

	code := ""

	for i := 0; i < len(lowerstringWord); i++ {
		code += morseCode[lowerstringWord[i]-97]
	}
	
	return code;
}

func bonus1() {

	compareValue := "-...-....-.--."

	var validWords []string;

	file, err := os.Open("./enable1.txt")
    if err != nil {
        log.Fatal(err)
    }
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    for scanner.Scan() {
		word := scanner.Text();
		if len(word) <= 14 || len(word) >= 4 {
			code := getMorse(word);
			if(code == compareValue) {
				validWords = append(validWords, word);
			}
		}
	}
	
	fmt.Printf("Valid words: %v", validWords);
}