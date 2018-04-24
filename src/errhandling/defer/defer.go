package main

import (
	"fmt"
	"os"
	"bufio"
	"functional"
)

/*
	defer 确保调用在函数结束时发生
	多个 defer ，先进后出
 */

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return
	//panic("error occurred")
	fmt.Println(4) //在 return 或 panic 之后，不会被输出
}

//defer 的使用场景
func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := functional.Fabonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	tryDefer()
	writeFile("fib.txt")
}
