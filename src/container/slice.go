package main

import "fmt"

/*
	go语言中的slice操作
	slice本身没有数据，是对底层数组的view
 */

func updateSlice(slice []int) {
	slice[0] = 100
}

func main() {

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("slice array arr[2:6]", arr[2:6])
	fmt.Println("slice array arr[2:]", arr[2:])
	fmt.Println("slice array arr[:6]", arr[:6])
	fmt.Println("slice array arr[:]", arr[:])

	s1 := arr[2:6]
	fmt.Println(s1)
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	s2 := s1[3:5]
	fmt.Println("s2 = ", s2)

	//slice具有扩展性，其三个属性: ptr(指针) len(长度) cap(容量)
	fmt.Printf("s1=%v, len=%d, cap=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len=%d, cap=%d\n", s2, len(s2), cap(s2))

	arr[2] = 2
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
	fmt.Println("append()")
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Printf("s3=%v, cap=%d\n", s3, cap(s3)) //未超过cap
	fmt.Printf("s4=%v, cap=%d\n", s4, cap(s4)) //超过之前的cap，系统重新分配
	fmt.Printf("s5=%v, cap=%d\n", s5, cap(s5))
	//当向 s2 append() 的时候，会替换掉之前位置上的7，因为此时还没有超过cap的值
	//当append超过cap的值时，系统会重新分配更大的底层数组
	fmt.Println("arr=", arr)
}
