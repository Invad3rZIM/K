package main

import "fmt"

type Node struct {
	left   *Node
	right  *Node
	token  string
	lexeme string
}

func (n *Node) print() {

	// if n.left != nil && n.right != nil {
	// 	fmt.Println(n.left, n.right)
	// }
	// if n.left == nil && n.right != nil {
	// 	fmt.Println(n.token, n.right)
	// }
	if n.left == nil && n.right == nil {
		fmt.Println(n.token, n.lexeme)
	}
	// if n.left != nil && n.right == nil {
	// 	fmt.Println(n.left, n.token)
	// }

	if n.left != nil {
		n.left.print()
	}
	if n.right != nil {
		n.right.print()
	}
}

func Treeify(words []string, tokens *map[string]string) *Node {
	/*
		1. Trim a single Paren mark
		2. build Left Node
		3. Build right node
	*/

	leftStart := 0
	rightEnd := len(words) - 1

	var leftChild *Node
	var rightChild *Node

	if len(words[0]) > 1 {
		token := (*tokens)[words[0]]
		words[0] = words[0][0 : len(words[0])-1]
		(*tokens)[words[0]] = token
	} else {
		leftStart++
	}

	if len(words[len(words)-1]) > 1 {
		token := (*tokens)[words[len(words)-1]]
		words[len(words)-1] = words[len(words)-1][0 : len(words[len(words)-1])-1]
		(*tokens)[words[len(words)-1]] = token
	} else {
		rightEnd--
	}

	//Building Left Child...

	if (*tokens)[words[leftStart]] == "ParenLeft" {
		pLevel := len(words[leftStart])

		i := leftStart + 1

		for pLevel > 0 {
			if (*tokens)[words[i]] == "ParenLeft" {
				pLevel += len(words[i])
			}
			if (*tokens)[words[i]] == "ParenRight" {
				pLevel -= len(words[i])
			}
			i++
		}

		//At this point, leftStart and i will both be ( and )
		leftChild = Treeify(words[leftStart:i], tokens)
	} else { //if the left child is a leaf
		leftChild = &Node{token: words[leftStart], lexeme: (*tokens)[words[leftStart]]}
	}

	//Building Right Child

	if (*tokens)[words[rightEnd]] == "ParenRight" {
		pLevel := len(words[rightEnd]) * -1

		i := rightEnd

		//find where the rightNode actually starts
		for pLevel < 0 {
			i--
			if (*tokens)[words[i]] == "ParenLeft" {
				pLevel += len(words[i])
			}
			if (*tokens)[words[i]] == "ParenRight" {
				pLevel -= len(words[i])
			}
		}

		//At this point, recurse into the right branch
		rightChild = Treeify(words[i:rightEnd+1], tokens)
	} else { //if the right child is a leaf...
		rightChild = &Node{token: words[rightEnd], lexeme: (*tokens)[words[rightEnd]]}
	}

	return &Node{left: leftChild, right: rightChild}
}

//Given an order and the tokens, returns parallel arrays of all the newline starts&ends
func GetNewLineIndices(order []string, tokens map[string]string) (*[]int, *[]int) {
	starts := []int{}
	ends := []int{}

	pLevel := 0

	for index, word := range order {
		if pLevel == 0 {
			starts = append(starts, index)
		}

		if tokens[word] == "ParenLeft" {
			pLevel = pLevel + len(word)
		}

		if tokens[word] == "ParenRight" {
			pLevel = pLevel - len(word)
		}

		if pLevel == 0 {
			ends = append(ends, index+1)
		}
	}

	return &starts, &ends
}
