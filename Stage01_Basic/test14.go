package main

import (
	"fmt"
)

func main() {
	//演示关系运算符的使用
	var n1 = 8
	var n2 = 6
	fmt.Println(n1 == n2) //false
	fmt.Println(n1 != n2) //true
	fmt.Println(n1 > n2)  //true
	fmt.Println(n1 >= n2) //true
	fmt.Println(n1 < n2)  //flase
	fmt.Println(n1 <= n2) //flase
	flag := n1 > n2
	fmt.Println("flag=", flag)
}
