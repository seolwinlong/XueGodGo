package main

import "fmt"

func main() {
	//switch分支控制语句
	var n1 int32 = 5
	var n2 int32 = 20

	switch n1 {
	case n2, 10, 5:
		fmt.Println("ok1")
	default:
		fmt.Println("没有匹配到")
	}
}
