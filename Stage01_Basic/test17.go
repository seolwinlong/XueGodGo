package main

import "fmt"

func main() {
	//单分支流程控制
	//编写一个程序，可以输入年龄，如果年龄大于18岁，则输出“你已经成年了”
	var age int
	fmt.Println("请输入年龄：")
	_, err := fmt.Scan(&age)
	if err != nil {
		return
	}
	if age >= 18 {
		fmt.Println("你已经成年了！")
	} else {
		fmt.Println("你还没成年！")
	}
}
