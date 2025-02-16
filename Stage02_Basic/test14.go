package main

import "fmt"

type PersonB struct {
	id   int
	name string
	age  int
}

type StudentB struct {
	PersonB
	score int
}

func (s *StudentB) coding() {
	//在值对象方法接收者修改不会影响源对象的信息
	s.name = "李雪娜"
	fmt.Println("正在编写Go语言面向对象代码")
}

func main() {
	stu := StudentB{PersonB{1002, "薛显皓", 5}, 100}
	fmt.Println(stu)
	stu.coding()
	fmt.Println(stu)
}

//在使用方法时，要注意如下几个问题：
//只要接收者类型不一样，这个方法就算同名，也是不同方法，不会出现重复定义函数的错误。
//如果接收者类型一样，但是方法的参数不一样，是会出现错误的。
//方法的接收者类型并非一定要是struct类型，type定义的类型别名、slice、map、channel、func类型等都可以。
//内置简单数据类型(int、float等)不行，可以定义类型别名来实现，interface类型不行。
//方法接收者是一个指针类型，则会自动解除引用。
