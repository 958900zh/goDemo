package tree

import "fmt"

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	fmt.Println(node.Value)
	node.Right.Traverse()
}

func (node *Node) TraverseFunc(f func(node *Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}