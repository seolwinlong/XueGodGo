package main

import "fmt"

//空接口（interface{}）不包含任何的方法，正因如此，所有的类型都实现了空接口，因此空接口可以存储任意类型的数值

func main() {
	//接口类型可以接收任意类型数据
	var v1 interface{} = 1              // 将int类型赋值给interface{}
	var v2 interface{} = "hello"        // 将string类型赋值给interface{}
	var v3 interface{} = 3.14           // 将float64类型赋值给interface{}
	var v4 interface{} = struct{}{}     //将结构体类型赋值给interface{}
	var v5 interface{} = []int{1, 2, 3} //将切片类型赋值给interface{}

	fmt.Println(v1, v2, v3, v4, v5)

	//接口类型不能直接进行转换和计算
	//a:=10
	//value:=a+v1

	//当函数可以接受任意的对象实例时，可以将其声明为interface{}，最典型的例子是标准库fmt中PrintXX系列的函数
	//func Printf(fmt string, args ...interface{})
	//func Println(args ...interface{})
}
