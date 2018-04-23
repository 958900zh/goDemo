package main

import "fmt"

func main() {

	var root tree.Node

	root = Tree.Node{Value: 5}
	root.left = &Node{}
	root.right = &Node{8, nil, nil}
	root.right.left = new(Node)
	fmt.Println(root)

	nodes := []Node{
		{value: 3},
		Node{},
		Node{9, nil, &root},
	}
	fmt.Println(nodes)
	root.getValue()
}
