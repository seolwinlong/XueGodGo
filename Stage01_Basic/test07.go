package main

import "fmt"

func main() {
	//基本数据类型
	//Go语言数据类型分为五种：布尔型，整型，浮点型，字符型，字符串型
	//定义bool类型变量
	//var 变量名 bool=true

	a := 10
	b := 20

	//比较两个数的大小，返回值为bool类型的结果
	fmt.Println(a > b)
	fmt.Println(a < b)

	//将表达式的结果赋值给bool类型变量c中
	var c bool = a < b
	fmt.Println(c)

	//格式化打印，使用占位符%t，输出一个bool类型数据
	fmt.Printf("%t", c)
}
