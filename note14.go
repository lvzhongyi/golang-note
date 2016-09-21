package main

import "fmt"
//流程控制 	goto语句
func main() {
	var i int = 1
	TAG:
	fmt.Println("这是tag:", i)
	goto TAG
}

