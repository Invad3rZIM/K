package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

const whiteSpace = 32
const doubleQuotes = 34
const dot = 46

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

	lineCount := 1

	for scanner.Scan() {
		t, l, err := parseLine(scanner.Text(), lineCount)

		if err != nil {
			return nil, err
		}

		tokens = append(tokens, t...)
		lexemes = append(lexemes, l...)
		lineCount++
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

func parseLine(str string, line int) ([]string, []string, error) {
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
		//state blank => we don't have a start index yet
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
					lexeme = "literal integer"
					state = "number"
				} else if r == doubleQuotes {
					start = i
					end = -1
					lexeme = "literal string"
					state = "inQuotes"
				} else if r == dot {
					start = i
					end = -1
					lexeme = "literal decimal"
					state = "decimal"
				} else {
					return nil, nil, parseError(line)
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
			if r == dot {
				state = "decimal"
				lexeme = "literal decimal"
				i++
				continue
			} else if !isNumber(r) {
				end = i
				words = append(words, string(runes[start:end]))
				lexemes = append(lexemes, lexeme)
				state = "blank"
				continue
			}
		} else if state == "decimal" {
			if r == dot {
				return nil, nil, parseError(line)
			}
			if !isNumber(r) {
				end = i

				if end == start+1 { //this invalidates a . decimal
					return nil, nil, parseError(line)
				}

				words = append(words, string(runes[start:end]))
				lexemes = append(lexemes, lexeme)
				state = "blank"
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
	//to handle the very last word
	if end == -1 {
		if lexeme == "literal decimal" && start+1 == i { //this invalidates a . decimal
			return nil, nil, parseError(line)
		}

		words = append(words, string(runes[start:i]))
		lexemes = append(lexemes, lexeme)
	}

	return words, lexemes, nil
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

func parseError(line int) error {
	return errors.New(fmt.Sprintf("Bad Parse on line: %d\n", line))
}
