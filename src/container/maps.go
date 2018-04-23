package main

import "fmt"

func main() {

	m1 := map[string]string{
		"name":    "zhangsan",
		"address": "beijing",
	}

	m2 := make(map[string]int) //type: empty map

	var m3 map[string]int //type: nil

	fmt.Println(m1, m2, m3)

	fmt.Println("*****Traversing map*****")
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	fmt.Println("*****Getting value*****")
	name, ok := m1["name"]
	fmt.Println(name, ok)

	//如果key不存在，会输出初始值(zero value) ok为false
	name, ok = m1["mame"]
	fmt.Println(name, ok)

	fmt.Println("*****Delete value*****")
	name, ok = m1["name"]
	fmt.Println("before delete : ", name, ok)
	delete(m1, "name")
	name, ok = m1["name"]
	fmt.Println("after delete : ", name, ok)

	/*
		在go语言中，map key的类型：
		除了 slice map function 之外
		Struct类型中不包含上述字段，也可以作为key
	 */
}
