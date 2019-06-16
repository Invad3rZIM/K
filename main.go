package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

const minus = 45
const whiteSpace = 32
const doubleQuotes = 34
const dot = 46

func main() {
	parseMap, err := readFile(os.Args[1])
	if err != nil {
		log.Panic(err)
	}

	parseMap.iterate()
}

func readFile(fileName string) (*ParseMap, error) {
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
		t, l, err := evaluateLine(scanner.Text(), lineCount)

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

	return NewParseMap(tokens, lexMap), nil
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

func isSymbol(ascii int) bool {
	return (ascii >= 33 && ascii <= 47) || (ascii >= 58 && ascii <= 64)
}
