package main

import "fmt"

//自定义类型

func main() {
	type x map[string]int
	x1 := x{"xx":1, "xxx":2}
	fmt.Println(x1)
}
