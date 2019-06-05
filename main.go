package main

//need to 1. read input
//2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	words, err := parseFile(os.Args[1])

	if err != nil {
		log.Panic(err)
	}

	for i := range words {
		fmt.Println(words[i])
	}
}

func parseFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var tokens []string

	for scanner.Scan() {
		tokens = append(tokens, parse(scanner.Text())...)
	}

	return tokens, nil
}

/*
	' ' => 32
	'(' => 40
	')' => 41
	'A' => 65
	'Z' => 90
	'a' => 97
	'z' => 122
	'0' => 48
	'9' => 57
*/

func parse(str string) []string {
	//standardize whitespacing
	runes := []rune(str)
	r := 0
	var words []string

	start, end := 0, 0
	state := 0

	i := 0

	for i < len(runes) {
		r = int(runes[i])
		//state 0 => we don't have a start index yet
		if state == 0 {
			if r != 32 {
				if r == 40 || r == 41 { // parenthesis
					words = append(words, string(runes[i]))
					i++
					state = 0
					continue
				} else if (r >= 65 && r <= 90) || (r >= 97 && r <= 122) { //letters
					start = i
					end = -1
					state = 1
				} else if r >= 48 && r <= 57 { //integers
					start = i
					end = -1
					state = 2
				}
			}
		} else if state == 1 { //start index maps to a letter
			if !((r >= 65 && r <= 90) || (r >= 97 && r <= 122)) {
				if r >= 48 && r <= 57 {
					state = 2
				} else {
					end = i
					words = append(words, string(runes[start:end]))
					state = 0
					continue
				}
			} else {
			}
		} else if state == 2 { //only numbers can be tokenized at this point
			if r < 48 || r > 57 {
				end = i
				words = append(words, string(runes[start:end]))
				state = 0
				continue
			}
		}

		i++
	}

	if end == -1 {
		words = append(words, string(runes[start:i]))
	}

	return words
}
