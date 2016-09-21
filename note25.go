package main

import "fmt"
//struct类型
func main() {
	//1
	var p Person
	p.name = "lvzhongyi"
	p.age = 23
	fmt.Println(p)

	//2
	p1 := Person{"lvzhongyi", 23}
	fmt.Println(p1)

	//3
	p2 := Person{age:23, name:"lvzhongyi"}
	fmt.Println(p2)

	//4
	p3 := new(Person)//这里返回的是一个指针
	fmt.Println(p3)
}

type Person struct {
	name string
	age  int
}