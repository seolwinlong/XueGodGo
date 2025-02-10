package main

import "fmt"

// 交换两个变量，但是不允许使用中间变量
func main() {
	a := 10
	b := 20
	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)
}
