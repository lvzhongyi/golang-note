package main

import "fmt"
//defer语句  （如果有return,在return之后执行）
//当函数执行到最后时，这些defer语句会按照逆序执行
// 如果有很多调用defer，那么defer是采用后进后出

func main() {
	test1()
}

func test1() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
