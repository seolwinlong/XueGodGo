package main

import (
	"fmt"
	"time"
)

func newTask2() {
	i := 0
	for {
		i++
		fmt.Printf("goroutine:%d\n", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go newTask2()
	fmt.Println("主程序退出")
}

//结果1
//main goroutine exit
//Process finished with exit code 0

//结果2
//main goroutine exit
//goroutine:1
//Process finished with exit code 0

//第一种结果分析为：当打印函数执行结束后，主函数的Goroutine退出，可是启动的另外一个Goroutine并没有执行完成就退出了。
//第二种结果分析为：当打印函数执行结束后，同时启动的另外一个Goroutine执行打印内容后进入延时操作后，主函数的Goroutine退出。
