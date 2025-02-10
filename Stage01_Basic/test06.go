package main

import "fmt"

func main() {
	//常用的标准输出函数有以下三种，都会将数据输出到标准设备，如屏幕
	//输出任意类型并换行
	fmt.Println("fmt.Println()输出的数据，输出结果换行")
	//输出任意类型数据，不换行
	fmt.Print("fmt.Print()输出的数据，输出结果不换行")
	//格式化输出数据，不换行
	fmt.Printf("fmt.Printf()输出的数据，格式化输出不换行 占位符：%s", "xuegod")

}
