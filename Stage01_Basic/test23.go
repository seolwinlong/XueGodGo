package main

import "fmt"

//形参与实参

// 形参示例代码
func param(num int) int {
	//num是形参，由函数内部定义。
	//调用函数时会把参数的值赋给num形参
	num++
	fmt.Println(num)
	return num
}

func param2(num *int) {
	//num是实参，是一个指针类型
	//调用函数时会把参数的内存地址赋值给num
	//通过指针修改变量的值会使外部变量一起发生改变
	*num++
	fmt.Println(*num)
	return
}

func main() {
	argv := 10
	//argv作为参数把变量的值传递给了param
	argv = param(argv)
	fmt.Println(argv)

	argv2 := 10
	param2(&argv2)
	fmt.Println(argv2)
}
