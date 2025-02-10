package main

import "fmt"

func main() {
	a := 10   //十进制
	b := 010  //八进制
	c := 0x10 //十六进制
	//使用占位符打印二进制
	fmt.Printf("%b\n", a)
	//使用占位符打印十进制
	fmt.Printf("%d\n", a)
	//使用占位符打印八进制
	fmt.Printf("%o\n", b)
	//使用占位符打印十六进制
	fmt.Printf("%x\n", c)
}
