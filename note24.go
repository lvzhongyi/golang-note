package main
// _  操作
// _操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数。
import (
	_ "./temp"
	"fmt"
)

func main() {
	fmt.Println("note24")
}

