package main

import "fmt"

// 切片
func main() {
	//切片定义的基本语法
	//生成一个空切片，示例：var arr1 []int
	//生成一个空数组，示例：var arr1 [10]int
	//总结：切片声明时不需要说明长度，[]没有声明长度，说明这是一个切片，而不是一个数组。
	//演示切片的基本使用
	var arr []int
	fmt.Println("空切片的值是：", arr)

	var intArr = [5]int{1, 11, 22, 55, 76} //定义一个数组
	//声明一个切片slice，并将数组intArr[1:3]的值赋给切片
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice=", slice)
	fmt.Println("slice的元素个数：", len(slice))
	fmt.Println("slice的容量：", cap(slice))
	slice[0] = 100
	fmt.Printf("slice[0]=%v,intArr[1]=%v\n", slice[0], intArr[1])
}
