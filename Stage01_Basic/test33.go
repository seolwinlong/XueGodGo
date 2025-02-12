package main

import "fmt"

// recover错误拦截
/*运行时panic异常一旦被引发就会导致程序崩溃。
Go语言提供了专用于“拦截”运行时panic的内建函数“recover”。
它可以是当前的程序从运行时panic的状态中恢复并重新获得流程控制权。*/

func Demo1(i int) {
	//定义10个元素的数组
	var arr [10]int

	//错误拦截要在产生错误前设置
	defer func() {
		//设置recover拦截错误信息
		err := recover()
		//产生panic异常 打印错误信息
		if err != nil {
			fmt.Println(err)
		}
	}()

	//根据函数参数为数组元素赋值
	//如果i的值超过数组下标，会报错：数组下标越界
	arr[i] = 10
}

func main() {
	Demo1(10)
	//产生错误后，程序继续执行
	fmt.Println("程序继续执行")
}
