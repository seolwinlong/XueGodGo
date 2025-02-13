package main

import "fmt"

// Student 结构体
// 结构体是由一系列具有不同类型的数据构成的数据集合。
// 注意：结构体成员在定义时不能赋值。
// 结构体定义
type Student struct {
	id    int
	name  string
	age   int
	sex   rune
	score float32
	addr  string
}

func main() {
	// 结构体初始化
	stu := Student{1001, "刘能", 47, '男', 100, "象牙山"}
	stu.id = 1002
	fmt.Println(stu)

	//结构体切片
	//定义结构体变量只能存储一条信息，如果使用结构体存储多个信息，需要使用结构体数组或切片。
	slice := []Student{
		{1001, "刘能", 47, '男', 100, "象牙山"},
		{1002, "谢广坤", 48, '男', 101, "象牙山"},
		{1003, "王大拿", 49, '男', 102, "象牙山"},
	}
	fmt.Println(slice)

	//切片扩容
	slice = append(slice, Student{1004, "王老七", 99, '男', 104, "象牙山"},
		Student{1004, "王老七", 99, '男', 104, "象牙山"})
	fmt.Println(slice)
	//遍历打印1
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	//遍历打印2
	for _, v := range slice {
		fmt.Println(v)
	}
	//切片截取
	s := slice[1:4]
	fmt.Println(s)
}
