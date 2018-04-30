package main

import (
	"fmt"
)

func doWork(i int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("worker %d recieved %c\n", i, n)
		done <- true
	}
}

func createWorker(i int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}

	go doWork(i, w.in, w.done)

	return w
}

type worker struct {
	in   chan int
	done chan bool
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		<-workers[i].done
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		<-workers[i].done
	}
}

func main() {
	chanDemo()

}
