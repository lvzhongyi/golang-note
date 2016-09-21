package main

import "fmt"
//map
func main() {
	//申明一个key为string,value为int的字典
	var a map[string]int
	//通过make关键字申明
	b := make(map[string]int)
	//a["woqu"] = 4
	fmt.Println(a)
	b["t"] = 5
	fmt.Println(b)

	c:=map[string]int{"aa":1,"bb":2}
	fmt.Println(c)

	zz,ok:=c["aa"]
	if ok {
		fmt.Println(zz)
	}

	delete(c,"bb")
	fmt.Println(c)
}
