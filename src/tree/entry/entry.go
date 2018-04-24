package main

import (
	"fmt"
	"tree"
)

/*
	配置GOPATH:
	在GOPATH下约定有三个目录： src pkg bin
	src存放源代码(比如：.go .c .h .s等)
	pkg编译时生成的中间文件（比如：.a）
	bin编译后生成的可执行文件（为了方便，可以把此目录加入到 $PATH 变量中，如果有多个gopath，那么使用${GOPATH//://bin:}/bin添加所有的bin目录）
 */

func main() {

	var root tree.Node

	root = tree.Node{Value: 5}
	root.Left = &tree.Node{}
	root.Left.Right = &tree.Node{Value: 6}
	root.Right = &tree.Node{8, nil, nil}
	root.Right.Left = new(tree.Node)
	fmt.Println(root)

	nodes := []tree.Node{
		{Value: 3},
		tree.Node{},
		tree.Node{9, nil, &root},
	}
	fmt.Println(nodes)
	root.GetValue()
	root.Traverse()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println(nodeCount)
}
