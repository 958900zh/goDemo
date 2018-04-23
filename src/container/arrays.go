package main

import "fmt"

/*
	在go语言中，所有的参数都是值传递
	数组也是值传递，这一点跟其他语言不同
 */
func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Print(i, "=", v, "\t")
	}
}

/*
	若想传引用，则可以使用指针
 */
func printArrayPtr(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Print(i, "=", v, "\t")
	}
}

func main() {

	//定义数组
	var a [5]int
	b := [5]int{1, 2, 3, 4, 5}
	c := [...]int{6, 7, 8}

	//二维数组
	d := [3][5]int{}

	//遍历数组
	for i := 0; i < len(b); i++ {
		fmt.Printf("%d\t", b[i])
	}

	//range有两个返回值，下标和值
	for i, v := range b {
		fmt.Print(i, "=", v, "  ")
	}

	fmt.Println(a, b, c, d)

	fmt.Println("func printArray()")
	printArray(b)
	fmt.Println(b)
	//printArray(c) 在go语言中，[3]int 和 [5]int 是两种不同的类型

	fmt.Println("func printArrayPtr()")
	printArrayPtr(&b)
	fmt.Println(b)
}
