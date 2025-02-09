package main

import "fmt"

// 常量的定义与使用
func main() {
	//常量集（iota枚举）
	//第一个iota等于0，每当iota在新的一行被使用时，它的值都会自动加1
	/*	const (
		a = iota
		b = iota
		c = iota
	)*/

	//与前一个例子相同，所以a=0,b=1,c=2可以简写如下形式：
	/*	const(
		a=iota
		b
		c
	)*/

	//在同一行iota的值相同
	/*	const (
			a    = iota
			b, c = iota, iota
			d, e
		)
		println(a, b, c, d, e)*/

	//常量集中的值可以自定义
	/*	const (
			a = 123
			b = true
			c = "hello"
		)
		fmt.Println(a, b, c)*/

	const (
		a    = iota
		b    = "go"
		c    = "mk"
		d, e = iota, iota
	)
	fmt.Println(a, b, c, d)
}
