package main

import "fmt"
//流程控制	for
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)

	for y := 10; y > 0; y-- {
		if y == 5 {
			continue
			//break
		}
		fmt.Println("y:", y)
	}
	//与标记号配合使用,这样就能一次跳出多个for循环了
	Z:for z := 10; z > 0; z-- {
		for t := 0; t < 10; t++ {
			if t == 5 {
				break Z
			}
			fmt.Println("t:", t)
		}
	}

	//配合range使用，可以读取slice和map中的数据
	m := make(map[string]int)
	m["xx"] = 123
	m["tt"] = 456

	for k, v := range m {
		fmt.Println(k, v)
	}

}
