package main

type ParseTree struct {
	root *Node
}

//Given a parsemap, construct the valid parse tree
func NewParseTree(pm *ParseMap) *ParseTree {
	//get parsetree line breaks
	starts, ends := GetNewLineIndices(pm.order, pm.tokens)
	nodeChain := []*Node{}

	//for each linebreak, construct that node
	for index, _ := range *starts {
		node := Treeify(pm.order[(*starts)[index]:(*ends)[index]], &(pm.tokens))
		nodeChain = append(nodeChain, node)
	}

	root := &Node{}
	curr := root

	//then chain all the nodes in reverse order to build the root

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
