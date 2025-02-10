package main

import "fmt"

// 标准输入输出
func main() {
	//Go语言中提供三种标准输入输出函数
	//扫描，必须所有参数都被填充后才结束扫描
	/*	fmt.Scan()
		//扫描，但是遇到换行就结束扫描
		fmt.Scanln()
		//格式化扫描，换行就结束
		fmt.Scanf()*/
	x, y := 0, 0
	//提供标准输入设备或去内容并赋值给变量x,y
	//注意这里x，y变量前面一定要加上“&”符号，表示获取内存单元的地址，然后才能存储
	_, err := fmt.Scan(&x, &y)
	if err != nil {
		return
	}
	fmt.Println(x)
	fmt.Println(y)
}
