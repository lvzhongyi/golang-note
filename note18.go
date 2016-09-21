package main

import "fmt"
//传值与传指针
//当我们传一个参数值到被调用函数里面时，实际上是传了这个值的一份copy，
// 当在被调用函数中修改参数值的时候，调用函数中相应实参不会发生任何变化，
// 因为数值变化只作用在copy上。
func main() {
	a := 1
	fmt.Println("a", a)
	b := test1(a)
	fmt.Println("a", a)
	fmt.Println("b", b)

	c := 1
	fmt.Println("c", c)
	d := test2(&c)
	fmt.Println("c", c)
	fmt.Println("d", d)
}

func test1(a int) int {
	a = a + 1
	return a
}

func test2(c *int) int {
	*c += 1
	return *c
}
