package main

import (
	"errors"
	"fmt"
)

/*
	panic:
	1.停止函数的运行
	2.一直向上返回，执行每一层的defer
	3.如果没有遇见recover，程序直接退出

	recover:
	1.仅在defer中调用
	2.可以获取panic的值
	3.如果无法处理，可重新panic
 */

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred: ", err)
		} else {
			panic(fmt.Sprintf("I don't know this error: %s", err))
		}
	}()

	a, b := 5, 0
	panic(123)
	fmt.Println(a / b)
	panic(errors.New("this is an error"))
}

func main() {
	tryRecover()
}
