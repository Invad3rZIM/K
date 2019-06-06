package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const whiteSpace = 32

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
	state := "looking"

	i := 0

	for i < len(runes) {
		r = int(runes[i])
		//state 0 => we don't have a start index yet
		if state == "looking" {
			if r != whiteSpace {
				if isParenthesis(r) { // parenthesis
					words = append(words, string(runes[i]))
					i++
					state = "looking"
					continue
				} else if isLetter(r) { //letters
					start = i
					end = -1
					state = "letter"
				} else if isNumber(r) { //integers
					start = i
					end = -1
					state = "number"
				}
			}
		} else if state == "letter" { //start index maps to a letter
			if !isLetter(r) {
				if isNumber(r) {
					state = "number"
				} else {
					end = i
					words = append(words, string(runes[start:end]))
					state = "looking"
					continue
				}
			} else {
			}
		} else if state == "number" { //only numbers can be tokenized at this point
			if !isNumber(r) {
				end = i
				words = append(words, string(runes[start:end]))
				state = "looking"
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

func isNumber(ascii int) bool {
	return ascii >= 48 && ascii <= 57
}

func isLetter(ascii int) bool {
	return (ascii >= 65 && ascii <= 90) || (ascii >= 97 && ascii <= 122)
}

func isParenthesis(ascii int) bool {
	return ascii == 40 || ascii == 41
}
