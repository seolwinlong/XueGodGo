package main

import "fmt"

// Go语言提供了两种精度的浮点数float32和float64
func main() {
	pi := 3.1415926
	//格式化打印 使用占位符%d 输出一个float类型数据
	//%f 默认保留六位小数 会四舍五入
	fmt.Printf("%f\n", pi)

	//设置保留两位小数
	fmt.Printf("%.2f\n", pi)
}
