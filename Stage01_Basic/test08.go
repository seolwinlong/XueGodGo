package main

import "fmt"

// Go语言的整数类型一共有10个
func main() {
	//有符号整型
	var m = -10
	//无符号整型
	var n uint = 10
	//格式化打印，使用占位符%d,输出一个int类型数据
	fmt.Printf("m=%d,n=%d\n", m, n)
	fmt.Printf("m的类型是：%T", m)
}
