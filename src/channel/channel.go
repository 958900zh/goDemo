package main

import (
	"fmt"
	"time"
)

func worker(i int, c chan int) {
	for {
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("worker %d recieved %c\n", i, n)
	}
}

func createWorker(i int) chan<- int {
	c := make(chan int)

	go worker(i, c)

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
	var channels [10] chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
}

func bufferedChannel() {
	c := make(chan int, 3) //设置 channel 的缓冲区大小为 3

	go worker(0, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)

	go worker(0, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)

	time.Sleep(time.Millisecond)
}

func main() {

	channelClose()
	time.Sleep(time.Second)
}
