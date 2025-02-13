package main

import "fmt"

func swap(a, b int) {
	//交换变量的值
	a, b = b, a
	fmt.Println("swap函数中值", a, b)
}

func swap2(a, b *int) {
	//指针间接交换变量的值
	*a, *b = *b, *a
	fmt.Println("swap函数中的值：", *a, *b)
}

// 指针
func main() {
	//可以使用运算符 & （取地址运算符）来获取数据的内存地址。
	i := 0
	//使用格式化打印变量的内存地址
	//%p是一个占位符，输出一个十六进制的地址格式
	fmt.Printf("%p\n", &i)

	//可以通过指针变量来存储，所谓的指针变量：就是用来存储任何一个值的内存地址。
	p := &i
	//打印指针变量p的值，同时也是i的地址
	fmt.Println(p)

	//在定义指针变量指向数据的内存地址后，可以通过指针间接操作数据。
	//需要使用运算符 *（取值运算符）操作数据。
	*p = 20
	fmt.Println(i)

	//在使用指针变量时，要注意定义指针默认值为nil。
	//如果直接操作指向nil的内存地址或报错。
	var p1 *int
	fmt.Println(p1)

	//*p1 = 123
	//fmt.Println(*p1)
	//panic: runtime error: invalid memory address or nil pointer dereference
	//所以，在使用指针变量时，一定要让指针变量有正确的指向。

	//new()函数
	//new( )函数的作用就是动态分配空间，不需要关心该空间的释放，Go语言会自动释放。
	var p2 *int
	//创建一个int大小的内存空间，返回值为*int
	p2 = new(int)
	fmt.Println(*p2)
	fmt.Println(p2)

	//指针做函数参数
	//普通变量作为函数参数进行传递是值传递
	a := 10
	b := 20
	swap(a, b)
	fmt.Println("main函数中的值：", a, b)
	//值传递 形参不能改变实参的值
	swap2(&a, &b)
	fmt.Println("main函数中的值", a, b)
}
