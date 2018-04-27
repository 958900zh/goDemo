package main

import (
	"fmt"
	"time"
)

func worker(i int, c chan int) {
	for {
		fmt.Printf("worker %d recieved %c\n", i, <-c)
	}
}

func createworker(i int) chan int {
	c := make(chan int)

	go func() {
		fmt.Printf("worker %d recieved %c\n", i, <-c)
	}()

	return c
}

func chanDemo() {
	//定义 channel 此时 c = nil
	//var c chan int
	c := make(chan int)
	go worker(0, c)
	c <- 0
	c <- 1
}

func chanDemo2() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
}

func chanDemo3() {
	var channels [10] chan int
	for i := 0; i < 10; i++ {
		channels[i] = createworker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
}

func main() {

	chanDemo3()
	time.Sleep(time.Second)
}
