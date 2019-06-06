package main

import (
	"bufio"
	"log"
	"os"
)

const whiteSpace = 32
const doubleQuotes = 34

func main() {
	parseMap, err := parseFile(os.Args[1])

	if err != nil {
		log.Panic(err)
	}

	parseMap.iterate()
}

func parseFile(fileName string) (*ParseMap, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var tokens []string
	var lexemes []string

	for scanner.Scan() {
		t, l := parseLine(scanner.Text())
		tokens = append(tokens, t...)
		lexemes = append(lexemes, l...)
	}

	lexMap := make(map[string]string)

	for i, _ := range tokens {
		lexMap[tokens[i]] = lexemes[i]
	}

	pm := ParseMap{
		order:   tokens,
		lexemes: lexMap,
	}

	return &pm, nil
}

func parseLine(str string) ([]string, []string) {
	//standardize whitespacing
	runes := []rune(str)
	var r int
	var words []string
	var lexemes []string
	var lexeme string
	state := "blank"
	start, end := 0, 0

	i := 0

	for i < len(runes) {
		r = int(runes[i])
		//state looking => we don't have a start index yet
		if state == "blank" {
			if r != whiteSpace {
				if isParenthesis(r) {
					lexeme = "separator"
					words = append(words, string(runes[i]))
					lexemes = append(lexemes, lexeme)
					i++
					continue
				} else if isLetter(r) {
					start = i
					end = -1
					lexeme = "identifier"
					state = "letter"
				} else if isNumber(r) {
					start = i
					end = -1
					lexeme = "literal"
					state = "number"
				} else if r == doubleQuotes {
					start = i
					end = -1
					lexeme = "literal"
					state = "inQuotes"
				}
			}
		} else if state == "letter" { //start index maps to a letter
			if !isLetter(r) && !isNumber(r) {
				end = i
				words = append(words, string(runes[start:end]))
				lexemes = append(lexemes, lexeme)
				state = "blank"
				continue
			} else {
			}
		} else if state == "number" { //only numbers can be tokenized at this point
			if !isNumber(r) {
				end = i
				words = append(words, string(runes[start:end]))
				lexemes = append(lexemes, lexeme)
				state = "blank"
				continue
			}
		} else if state == "inQuotes" { //handles quoted strings
			if r == doubleQuotes {
				end = i
				words = append(words, string(runes[start:end+1]))
				lexemes = append(lexemes, lexeme)
				state = "blank"
			}
		}

		i++
	}
	//to append the very last word
	if end == -1 {
		words = append(words, string(runes[start:i]))
		lexemes = append(lexemes, lexeme)
	}

	return words, lexemes
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
