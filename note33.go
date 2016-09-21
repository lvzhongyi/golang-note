package main

import (
	"strconv"
	"fmt"
)
//使用Comma-ok（断言）或者switch来判断interface中保存的数据的类型

// Comma-ok判断,语法    $value$,$ok$:= $变量$.($type$) ----   value>变量的值    ok>判断变量是否是&type&类型
type Element interface{}

type List []Element

type Person struct {
	name string
	age  int
}
//定义了string方法，实现了fmt.Stringer
func (p Person) String() string {
	return "(name:" + p.name + "- age:" + strconv.Itoa(p.age) + "years)"
}

func test1() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "hello"
	list[2] = Person{"lvzhongyi", 17}
	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is an string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is an Person and its value is %s", index, value)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}
}

//*****************************************

//switch判断
func test2() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "hello"
	list[2] = Person{"lvzhongyi", 17}
	/**
	这里有一点需要强调的是:$变量$.(type) 语法不能在switch外的任何逻辑里面使用，如果
	你要在switch外面判断一个类型就使用Comma-ok
	 */

	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is an string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is an Person and its value is %s", index, value)
		default:
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}
}
func main() {
	test1()
	test2()

}
