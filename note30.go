package main

import "fmt"

//method的继承


type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human)SayHi() {
	fmt.Printf("Hi,I am %s you can call me on %s\n", h.name, h.phone)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi,I am %s you can call me on %s,I work at %s\n", e.name, e.phone, e.company)
}

func main() {
	mark := Student{Human{"Mark", 25, "1111234"}, "Harvard"}
	sam := Employee{Human{"sam", 22, "32321122"}, "Golang inc"}
	mark.SayHi()
	sam.SayHi()
}
