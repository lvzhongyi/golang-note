package main

import "fmt"

//切片slice
func main() {
	// 和声明array一样，只是少了长度
	a := []int{1, 3, 4, 5, 6}
	fmt.Println(a)

	//slice可以从一个数组或一个已经存在的slice中再次声明。
	// slice通过array[i:j]来获取，其中i是数组的开始位置，
	// j是结束位置，但不包含array[j]，它的长度是j-i

	var b = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	fmt.Println(b)

	c := b[2:5]
	fmt.Println(c)

	d := b[3:5]
	fmt.Println(d)

	e := b[3:]
	fmt.Println(e)

	f := b[:5]
	fmt.Println(f)

	g := b[:]
	fmt.Println(g)

	//slice是引用类型，所以当引用改变其中元素的值时，其它的所有引用都会改变该值


	fmt.Println("长度:", len(e))
	fmt.Println("最大容量是:", cap(e))//容量和长度只有在map或者slice中有其它对象的时候才会有区别，这句话可能有问题，自己的理解
	//追加元素
	fmt.Println(append(e, 'z'))

	//go1.2以后，slice接收三个参数，第三个参数用来指定容量,容量=第三个数-第一个数
	var x1 [10]int
	//如果没有第三个参数，容量就由第一个参数和被切片的数组长度决定 10-3
	x2 := x1[3:6:7] //容量=7-2
	fmt.Println(cap(x2))
}
