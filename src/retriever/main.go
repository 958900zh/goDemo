package main

import (
	"fmt"
	"retriever/mock"
	"retriever/real"
	"time"
)

/*
	go语言中常用的接口：
	stringer
	reader
	writer
 */

//定义接口
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "http://www.baidu.com"

func download(r Retriever) string {
	return r.Get(url)
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{"contents": "another fake retriever"})

	return s.Get(url)
}

func main() {
	var r Retriever
	r = &mock.Retriever{Contents: "this is a fake http"}
	fmt.Printf("%T %v\n", r, r) //%T 类型， %v 值
	fmt.Println("implement stringer => ", r)
	r = real.Retriever{
		UserAgent: "Jack",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T %v\n", r, r)
	fmt.Println("not implement stringer => ", r)
	//fmt.Print(download(r))

	retriever := &mock.Retriever{Contents: "this is a fake http"}
	fmt.Println(session(retriever))
}
