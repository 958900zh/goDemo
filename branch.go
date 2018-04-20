package main

import (
	"io/ioutil"
	"fmt"
)

func ifDemo() {
	const filename = "abc.txt"
	/* 写法一
	contents, err :=ioutil.ReadFile(filename)

	if error != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}*/

	//写法二
	/*
		if的条件里可以赋值
		被赋值的变量的作用域就在这个if语句里
	 */
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

/*
	go语言中，case会自动break
 */
func switchDemo(a, b int, op string) int {
	var result int

	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		//panic 报错并中断程序执行
		panic("unsupported operator " + op)
	}

	return result
}

func main() {

	ifDemo()
	fmt.Println(switchDemo(1, 2, "+"))
}
