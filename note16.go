package main

import "fmt"
//switch case
func main() {
	expr := 2
	switch expr {
	case 0:
		fmt.Println(0)
	case 1, 2, 3:
		fmt.Println(1)
	case 4:
		fmt.Println(2)
	default:
		fmt.Println(nil)
	}

	//加上fallthrough强制执行后面的case代码
	switch expr {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
		fallthrough
	case 3:
		fmt.Println(3)
		fallthrough
	default:
		fmt.Println(nil)
	}
}
