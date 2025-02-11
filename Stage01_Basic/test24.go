package main

import "fmt"

//不定参函数

//第一种普通fori方式
//func sum(args ...int) int {
//	value := 0
//	for i := 0; i < len(args); i++ {
//		value += args[i]
//	}
//	return value
//}

//第二种forr方式

func sum(args ...int) int {
	value := 0
	for _, arg := range args {
		value += arg
	}
	return value
}

//函数类型：在Go中，函数也可以被当成数据类型
//定义函数类型后，可以将其当作数据类型或者函数参数类型，通过type定义，可以理解为函数别名

// 定义函数类型
type functype func(a, b int) int

func add(a, b int) int {
	return a + b
}

func main() {
	//在定义函数的时候根据需求指定参数的个数和类型，但是有时候无法确定参数的个数呢
	//例如：计算n个整数的和
	//在数据集合后面跟三个点，就是表示该参数可以接受0或者多个整数值
	//计算n个整数的和s
	result := sum(1, 2, 3, 4, 5)
	fmt.Println(result)

	//通过函数类型定义变量f并指向add函数
	var f functype = add
	//自动推导类型
	//f := add
	value := f(10, 20)
	fmt.Println(value)
}

//函数返回值注意事项：
//1、保证return语句中表达式的值和函数返回数据类型时同一类型
//2、如果函数返回的类型和return语句中表达式的值不一致，程序会报错
//3、return语句的另一个作用作为中断return所在的执行函数
//4、如果函数带返回值，return后面必须跟着值
