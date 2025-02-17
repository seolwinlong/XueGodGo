package main

import (
	"fmt"
	"runtime"
	"time"
)

//调用 runtime.Goexit() 将立即终止当前 goroutine 执⾏，调度器确保所有已注册 defer延迟调用被执行。

func main() {
	go func() {
		defer fmt.Println("A defer")
		func() {
			defer fmt.Println("B defer")
			runtime.Goexit() //终止当前goroutine
			fmt.Println("B") //不会执行
		}()
		fmt.Println("A") //不会执行
	}() //匿名函数调用格式
	//目的不让主goroutine提前结束
	time.Sleep(5 * time.Second)
}
