package main

import "fmt"

//自定义类型中的method
const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

//Color作为byte的别名
type Color byte

//定义一个struct Box，用来描述盒子的属性
type Box struct {
	width  float64
	height float64
	depth  float64
	color  Color
}
//定义一个Box类型的切片slice
type BoxList []Box

/**
以下为method
 */

//求体积,方法 receiver Box
func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}
//设置颜色，方法 receiver Box
func (b *Box) SetColor(c Color) {
	//b.color = c
}

//最大的盒子是什么颜色,receiver BoxList
func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

//将所有盒子变成黑色,receiver BoxList
func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
		//(&(bl[i])).SetColor(BLACK)//也可以这么写
	}
}
//根据盒子颜色，得到对应的颜色名称 ,receiver Color
func (c Color) toString() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}
	fmt.Printf("we have %d boxes in our set\n", len(boxes))

	fmt.Println("第一个盒子的体积是", boxes[0].Volume(), "cm³")
	fmt.Println("最后一个盒子的颜色是：", boxes[len(boxes) - 1].color.toString())

	fmt.Println("体积最大的一个盒子颜色是:", boxes.BiggestColor().toString())

	boxes.PaintItBlack()//将所有盒子置为黑色
	fmt.Println("The color of the second one is", boxes[1].color.toString())
	///////////////////
	//这里需要注意的是，我之前认为boxes.PaintItBlack传入的是copy,最终打印不会修改，但是我忘记了一点（slice是引用类型）
	//			还有一点是：setColor()方法接收的是一个指针
	///////////////////
	fmt.Println(boxes)
}
