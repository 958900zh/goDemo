package main

import (
	"fmt"
	"time"
	"runtime"
)

/*
	协程:
	轻量级“线程”
	非抢占式多任务处理，由协程主动交出控制权

	go routine 可能切换的点:
	I/O select
	channel
	Gosched
	等待锁
	函数调用（有时）
 */

func main() {

	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("print value %d\n", i)
				runtime.Gosched()
			}
		}(i)
	}

	time.Sleep(time.Second)
}
