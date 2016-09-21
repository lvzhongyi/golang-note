package main

import (
	"errors"
	"fmt"
)
//错误类型
func main() {
	err:=errors.New("这是一个错误，次奥")
	fmt.Print(err)
}
