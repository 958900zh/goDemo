package tree

import "fmt"

/*
	在go语言中，没有继承和多态，只有结构 struct
 */
type Node struct {
	Value       int
	Left, Right *Node
}

//为结构定义方法
func (node Node) GetValue() {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil " +
			"node. Ignored.")
		return
	}
	node.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}
