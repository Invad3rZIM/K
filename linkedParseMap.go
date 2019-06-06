package main

import "fmt"

type ParseMap struct {
	order   []string
	lexemes map[string]string
}

func (pm *ParseMap) iterate() {
	for _, v := range pm.order {
		fmt.Printf("%s - %s\n", v, pm.lexemes[v])
	}
}
