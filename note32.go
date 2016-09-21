package main

import "fmt"

//空 interface
//不包括任何的method
/**
空interface对于描述起不到任何的作用(因为它不包含任何的method），
 但是空interface在我们需要存储任意类型的数值的时候相当有用，
 因为它可以存储任意类型的数值。它有点类似于C语言的void*类型
*/
func main() {
	var a interface{}
	var i int = 5
	s := "Hello word"
	a = i
	a = s
	fmt.Println(a)//这个函数的源码里面就是对空interface的一个灵活的使用

}
