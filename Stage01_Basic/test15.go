package main

import "fmt"

func main() {
	//逻辑运算符
	var a = true
	var b = false

	if a || b {
		fmt.Println("第二行 - 条件为 true")
	}
	/* 修改 a 和 b 的值 */
	a = false
	b = true
	if a && b {
		fmt.Println("第三行 - 条件为 true")
	} else {
		fmt.Println("第三行 - 条件为 false")
	}
	if !(a && b) {
		fmt.Println("第四行 - 条件为 true")
	}
}
