package main

import "fmt"
//流程控制  if
func main() {
	var i int = 5
	if i > 10 {
		fmt.Println("i大于10")
	} else {
		fmt.Println("i小于等于10")
	}
	if x := test(); x > 5 {
		fmt.Println("x大于5")
	} else {
		fmt.Println("x小于等于5")
	}
	//这一句报错，x生命周期在if代码块
	fmt.Println(x)
}

func test() int {
	return 5
}
