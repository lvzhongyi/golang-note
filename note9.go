package main

import "fmt"
//Array数组
func main() {
	var arr [10]int
	arr[0] = 42
	arr[1] = 12
	fmt.Println(arr)

	a := [10]int{1, 3, 4}
	fmt.Println(a)

	b := [...]int{4, 5, 6}
	fmt.Println(b)

	//二维数组，先x后y
	c := [4][5]int{{1, 2, 3}, {11, 22, 33}}
	fmt.Println(c)
}

