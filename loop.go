package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func convertToBin(n int) string {

	result := ""
	for ; n > 0; n /= 2 {
		result = strconv.Itoa(n%2) + result
	}

	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	/*
		go语言中没有while
	 */
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	/*
		go语言中死循环的写法:
		for {
			...
		}
	 */
}

func main() {

	fmt.Println(convertToBin(5), convertToBin(12))
	printFile("abc.txt")

}
