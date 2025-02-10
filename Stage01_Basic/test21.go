package main

import "fmt"

/*
return 使用在方法或者函数中，表示跳出所在的方法或函数，在讲解函数的时候，会详细的介绍。
 1. 如果 return 是在普通的函数，则表示跳出该函数，即不再执行函数中 return 后面代码，也可以理解成终止函数。
 2. 如果 return 是在 main 函数，表示终止 main 函数，也就是说终止程序。
*/
func main() {
	n := 30
	//演示return使用
	fmt.Println("ok1")
	if n > 20 {
		return
	}
	fmt.Println("ok2")
}
