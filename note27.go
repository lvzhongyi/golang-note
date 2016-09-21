package main

import "fmt"

//面向对象
//method  方法名前加上属于的struct 例：fun (t type) xxx(),xxx方法就属于type类型

//method除了可以定义在struct,还可以定义在任何自定义的类型，内置类型,struct等各种类型上


//假设有这么一个场景，你定义了一个struct叫做长方形，你现在想要计算他的面积，那么按照我们一般的思路应该会用下面的方式来实现
func main() {
	fmt.Println(Rectangle{20, 20}.area())
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}