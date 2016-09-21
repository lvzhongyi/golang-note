package main

import "fmt"

//struct类型的匿名字段

//当匿名字段是一个struct的时候，那么这个struct所拥有的全部字段都被隐式地引入了当前定义的这个struct。(相当于java接口继承)

func main() {
	t := t2{T1{1, 2}, 3, 4}
	fmt.Println(t.a, t.b, t.c, t.d)

	t.T1 = T1{11, 22}
	fmt.Println(t)

	t3 := t3{t2{T1{11, 22}, 3, 4}, 5, 6}
	fmt.Println(t3)

	t3.int = 5555
	fmt.Println(t3)


	t5:=t5{t4{"aa","bb"},33,"dd"}
	fmt.Println(t5)
	fmt.Println(t5.phone,t5.t4.phone)
}

type T1 struct {
	a int
	b int
}
type t2 struct {
	T1
	c int
	d int
}

type t3 struct {
	t2
	int
	e int
}

type t4 struct {
	name  string
	phone string
}
type t5 struct {
	t4
	age   int
	phone string
}
