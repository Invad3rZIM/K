package main

type ParseTree struct {
	root *Node
}

func NewParseTree(pm *ParseMap) *ParseTree {
	starts, ends := GetNewLineIndices(pm.order, pm.tokens)
	nodeChain := []*Node{}

	for index, _ := range *starts {
		node := Treeify(pm.order[(*starts)[index]:(*ends)[index]], &(pm.tokens))
		nodeChain = append(nodeChain, node)
	}

	root := &Node{}
	curr := root

	index := len(nodeChain) - 1

	for index >= 1 {
		curr.right = nodeChain[index]
		curr.left = &Node{}

		curr = curr.left
		index--
	}

	curr.left = nodeChain[0]

	return &ParseTree{root: root}
}
