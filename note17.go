package main

import "fmt"
//多返回值方法
func main() {
	sum, product := test1(3, 4)
	fmt.Println(sum, product)
	sum, product = test2(3, 4)
	fmt.Println(sum, product)
	test3(5,5,5,5,6)
}
func test1(a, b int) (sum, product int) {
	sum = a + b
	product = a * b
	return
}

func test2(a, b int) (int, int) {
	return a + b, a * b
}
//变参函数(接收不定个数的参数)
func test3(arg ...int) {
	for _, b := range arg {
		fmt.Println(b)
	}
}