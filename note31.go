package main

import "fmt"

//interface
//概括：interface是一组method的组合，我们通过interface来定义对象的一组行为。
/**
Go语言采用的是“非侵入式接口",Go语言的接口有其独到之处：只要类型T的公开方法
完全满足接口I的要求，就可以把类型T的对象用在需要接口I的地方，所谓类型T的公
开方法完全满足接口I的要求，也即是类型T实现了接口I所规定的一组成员。这种做法的
学名叫做Structural Typing，有人也把它看作是一种静态的Duck Typing。
 */

/**
golang的接口思想：当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子
 */

//定义interface
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}
type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) SayHi() {
	fmt.Printf("Hi,I am %s you can call me on %s\n", h.name, h.phone)

}
func (h Human) Sing(lyrics string) {
	fmt.Println("la la la...", lyrics)
}
func (h Human)Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle...", beerStein)
}

func main() {
	var human Men
	human = Human{"lvzhongyi", 23, "17002023991"}
	human.SayHi()
	human.Sing("xxxx")
	human.Guzzle("tttt")
}
