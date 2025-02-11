package main

import "fmt"

//Go语言匿名函数
//Go语言支持匿名函数，匿名函数就是没有名字的函数，如果我们某个函数只希望使用一次，可以考虑使用匿名函数，匿名函数也可以实现多次调用

var (
	// Fun1 如果将匿名函数赋值给一个全局变量，那么这个匿名函数，就成为一个全局匿名函数
	Fun1 = func(n1, n2 int) int {
		return n1 * n2
	}
)

func main() {
	result := Fun1(2, 2)
	fmt.Println(result)

	//匿名函数
	func() {
		fmt.Println("Hello World")
	}()

	//匿名函数也可以传递参数
	func(a, b int) {
		fmt.Println(a + b)
	}(1, 3)

	//在定义匿名函数时就直接调用，这种方式匿名函数只能调用一次

	//案例演示，求两个数的和，使用匿名函数完成
	func(c, d int) {
		fmt.Println(c + d)
	}(5, 7)
}
