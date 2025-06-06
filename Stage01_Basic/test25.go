package main

import "fmt"

//Go语言递归函数
//递归概述：在调用一个函数的过程中又直接或间接地调用该函数本身

func test(n int) {
	if n > 2 {
		n--
		test(n) //递归必须向退出条件逼近，否则就是无限循环调用
	}
	fmt.Println("n=", n)
}

func main() {
	test(4)
}

//递归总结：
//1、函数地局部变量是独立地，不会相互影响
//2、递归必须向退出条件逼近，否则就是无限递归
//3、当一个函数执行完毕，或者遇到return，就会返回，遵守谁调用，就将结果返回给谁，同时当函数执行完毕或返回时，该函数本身也会被系统销毁
