package main

import (
	"fmt"
	"strings"
	"time"
	"bufio"
	"os"
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

func getListOfWordsFromFile(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func isPerfectBalance(input string) bool {
	dot := 0
	dash := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '.' {
			dot++
		} else {
			dash++
		}
	}
	return dot == dash
}