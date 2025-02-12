package main

import "fmt"

/*init 函数作用：
1、对变量进行初始化
2、检查/修复程序的状态
3、注册
4、运行一次计算（sync.Once）*/

// init优于主函数执行
func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}
func main() {
	//init() 不可以在函数中调用init
	fmt.Println("主函数执行")
}
