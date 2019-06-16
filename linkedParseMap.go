package main

import "fmt"

type ParseMap struct {
	order   []string
	lexemes map[string]string
	tokens  map[string]string
}

func (pm *ParseMap) iterate() {
	for _, v := range pm.order {
		fmt.Printf("%s %s %s\n", v, pm.lexemes[v], pm.tokens[v])
	}
}

func NewParseMap(order []string, lexemes map[string]string) *ParseMap {
	pm := ParseMap{order: order, lexemes: lexemes, tokens: make(map[string]string)}
	pm.tokenize()

	return &pm
}

func (pm *ParseMap) tokenize() {
	for _, word := range pm.order {
		var token string

		switch pm.lexemes[word] {
		case "identifier":
			if word == "let" {
				token = "IDBind"
			} else if word == "true" || word == "false" {
				token = "IDBoo"
			} else {
				token = "IDFree"
			}
			break
		case "operator":
			switch word {
			case ".":
				token = "OpFunc"
				break
			case "+":
				token = "OpSum"
				break
			case "-":
				token = "OpDiff"
				break
			case "*":
				token = "OpProd"
				break
			case "/":
				token = "OpRatio"
				break
			case "%":
				token = "OpMod"
				break
			case "=":
				token = "OpEq"
				break
			case "&":
				token = "OpAnd"
				break
			case "|":
				token = "OpOr"
				break
			case "^":
				token = "OpXor"
				break
			case "!":
				token = "OpNeg"
				break
			}
			break

		case "lsep":
			token = "ParenLeft"
			break
		case "rsep":
			token = "ParenRight"
			break
		case "literal int":
			token = "NumInt"
			break
		case "literal float":
			token = "NumFloat"
			break
		}

		pm.tokens[word] = token
	}
}
