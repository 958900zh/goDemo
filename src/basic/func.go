package main

import "fmt"

/*
	go语言中的函数可以有多个返回值
	返回值类型写在最后面
	函数可以作为参数
 */
func div(m int, n int) (int, int) {
	return m / n, m % n
}

/*
	go语言中的指针操作
 */
func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {

	a, b := div(28, 5)
	fmt.Println(a, b)

	//如果只想使用其中的一个返回值，另一个不想使用的返回值用 _ 来代替
	c, _ := div(16, 3)
	fmt.Println(c)

	m, n := 5, 8
	swap(&m, &n)
	fmt.Println(m, n)
}
