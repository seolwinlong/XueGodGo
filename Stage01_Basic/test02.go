package main

import "fmt"

// Go语言变量和常见的数据类型
func main() {
	//第一种：指定变量类型，声明后若不赋值，使用默认值
	//int的默认值是0，其他数据类型的默认值马上介绍
	var i int
	fmt.Printf("%d\n", i)

	//变量i赋值
	i = 6
	fmt.Println("i=", i)

	//第二种：根据值自行判定变量类型（类型推导）
	var number = 6
	fmt.Println("number=", number)

	//第三种：省略var关键词声明变量的方法，变量名：=“变量值”
	name := "jack"
	fmt.Println("name=", name)
}
