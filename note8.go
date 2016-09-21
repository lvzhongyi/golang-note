package main

import "fmt"
//iota，枚举申明
const (
	x = iota
	y = iota
	z = iota
	// 常量声明省略值时，默认和之前一个值的字面相同。

)
const (
	a1 = iota
	b2
	c3
)
const (
	a = iota
	b = "B"
	c = iota    //c=2
	d, e, f = iota, iota, iota //d=3,e=3,f=3
	_, _, g  //g = 4
)

func main() {
	fmt.Println(b)
}
