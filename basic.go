package main

import "fmt"

/*
	定义包变量
 */
var(
	aa int = 5
	bb bool = true
	ss string = "def"
)

/*
	go语言定义变量，先写变量名称，再写变量类型
 */
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

/*
	go语言支持类型推断
 */
func variableTypeDeduction() {
	var a, b, c, d = 5, 6, true, "abc"
	fmt.Println(a, b, c, d)
}

/*
	在go语言中 := 表示声明变量并赋值（推荐）
 */
func variableShorter() {
	a, b, c, d := 5, 6, true, "abc"
	fmt.Println(a, b, c, d)
}

/*
	go语言的常量定义
	在其他的语言中，常量的变量名一般都大写
	在go语言中，变量名大写是有其他含义的
 */
func consts() {
	const filename = "abc.txt"
	const a, b = 1, true
	fmt.Println(filename, a, b)
}

/*
	go语言中的枚举类型
 */
func enums() {
	/*方法一
	const(
		c = 0
		java = 1
		python = 2
		golang = 3
	)*/
	//方法二，使用 iota
	const (
		c      = iota
		java
		python
		golang
	)

	const (
		b  = 1 << (10 * iota)
		kb
		mb
		gb
	)

	fmt.Println(c, java, python, golang)
	fmt.Println(b, kb, mb, gb)
}

/*
	go语言的内建变量类型：
	bool string
	(u)int (u)int8 (u)int16 (u)int32 (u)int64 uintptr(指针)
	byte rune(类似于char)
	float32 float64 complex64 complex128
 */

func main() {

	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, ss)
	consts()
	enums()
}
