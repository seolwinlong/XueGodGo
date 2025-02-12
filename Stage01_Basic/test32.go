package main

import "fmt"

//defer延迟调用
//defer语句被用于预定对一个函数的调用。可以把这类被defer语句调用的函数称为延迟函数。
/*defer作用：
1、释放占用的资源
2、捕捉处理异常
3、输出日志*/

func Demo() {
	defer fmt.Println("学神教育")
	defer fmt.Println("在线教学")
	defer fmt.Println("月薪过万")
	defer fmt.Println("轻松就业")
}

func main() {
	Demo()
	//如果一个函数中有多个defer语句，它们会以LIFO（后进先出）的顺序执行
}
