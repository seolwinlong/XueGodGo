package main

import "fmt"

//当遇到不可恢复的错误状态的时候，如数组访问越界、空指针引用等，这些运行时错误会引起panic异常。
//一般而言，当panic异常发生时，程序会中断运行。随后，程序崩溃并输出日志信息。日志信息包括
//panic value和函数调用的堆栈跟踪信息。

func main() {
	fmt.Println("hello world")
	panic("调用panic函数，程序崩溃并输出异常")
	fmt.Println("学神教育强不强")
}
