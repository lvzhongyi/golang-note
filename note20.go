package main

import "fmt"
//函数作为值、类型
func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	odd := filter(slice, isOdd)//将函数作为值传递过去
	fmt.Println(odd)
	even := filter(slice, isEven) //将函数作为值传递过去
	fmt.Println(even)
}

func isOdd(number int) bool {
	if number % 2 == 0 {
		return false
	}
	return true
}
func isEven(number int) bool {
	if number % 2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, testInt testInt) []int {
	var result []int
	for _, value := range slice {
		if testInt(value) {
			result = append(result, value)
		}
	}
	return result
}

type testInt func(int) bool//声明一个函数类型