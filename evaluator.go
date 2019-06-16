package main

var operators map[string]bool

//Load keywords
func init() {
	ops := []string{"+", "-", "*", "/", "P", "C", "!", ".", "%", "=", "&", "|", "^"}
	operators = make(map[string]bool)

	for _, o := range ops {
		operators[o] = true
	}
}

//FSA used to parse the text for tokens and appropriate lexemes
func evaluateLine(str string, line int) ([]string, []string, error) {
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
				start = i
				end = -1

				if isParenthesis(r) {
					state = "parenthesis"
					if r == 40 {
						lexeme = "lsep"
					} else {
						lexeme = "rsep"
					}
				} else if r == doubleQuotes {
					state = "inQuotes"
					lexeme = "literal string"
				} else if isLetter(r) {
					state = "letter"
					lexeme = "identifier"
				} else if r == minus {
					state = "sub"
					lexeme = "sub"
				} else if isSymbol(r) {
					state = "symbol"
					lexeme = "operator"
				} else if isNumber(r) {
					state = "number"
					lexeme = "literal int"
				} else if r == dot {
					state = "decimal"
					lexeme = "literal float"
				}
			}
		} else if state == "parenthesis" {
			if (lexeme == "lsep" && r != 40) || (lexeme == "rsep" && r != 41) {
				end = i
				words = append(words, string(runes[start:i]))
				lexemes = append(lexemes, lexeme)
				state = "blank"
				continue
			}
		} else if state == "letter" {
			if r == whiteSpace {
				end = i
				words = append(words, string(runes[start:i]))
				lexemes = append(lexemes, lexeme)
				state = "blank"
			} else if !isNumber(r) && !isLetter(r) {
				return nil, nil, parseError(line)
			}
		} else if state == "number" {
			if r == dot {
				state = "decimal"
				lexeme = "literal float"
			} else if isParenthesis(r) || r == whiteSpace {
				end = i
				words = append(words, string(runes[start:i]))
				lexemes = append(lexemes, lexeme)
				state = "blank"
				continue
			} else if !isNumber(r) {
				return nil, nil, parseError(line)
			}
		} else if state == "inQuotes" {
			if r == doubleQuotes {
				end = i
				words = append(words, string(runes[start:i+1]))
				lexemes = append(lexemes, lexeme)
				state = "blank"
			}
		} else if state == "sub" {
			if isNumber(r) {
				lexeme = "literal int"
				state = "number"
			} else if isSymbol(r) {
				lexeme = "operator"
				state = "symbol"
			} else if r == whiteSpace {
				lexeme = "operator"
				state = "blank"
				words = append(words, string(runes[start:i]))
				lexemes = append(lexemes, lexeme)
				continue
			}
		} else if state == "symbol" {
			if isParenthesis(r) || r == whiteSpace {
				end = i
				words = append(words, string(runes[start:i]))
				lexemes = append(lexemes, lexeme)
				state = "blank"
				continue
			} else if r == doubleQuotes || !isSymbol(r) {
				return nil, nil, parseError(line)
			}
		} else if state == "decimal" {
			if isParenthesis(r) || r == whiteSpace {
				if start+1 == i {
					lexeme = "operator"
				}
				end = i
				words = append(words, string(runes[start:i]))
				lexemes = append(lexemes, lexeme)
				state = "blank"
				continue
			} else if !isNumber(r) {
				return nil, nil, parseError(line)
			}
		}

		i++
	}
	//to handle the very last word
	if end == -1 {
		if lexeme == "decimal" && start+1 == i { //this invalidates a . decimal
			lexeme = "operator"
		}

		words = append(words, string(runes[start:i]))
		lexemes = append(lexemes, lexeme)
	}

	for i, w := range words {
		if operators[w] {
			lexemes[i] = "operator"
		}

		if w == "true" || w == "false" {
			lexemes[i] = "literal bool"
		}
	}

	return words, lexemes, nil
}
