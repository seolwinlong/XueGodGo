package main

import "fmt"

// Go语言的字符有以下两种
// 一种是uint8类型或者叫byte类型，代表了ASCII码的一个字符
// 一种是rune类型，代表一个UTF-8字符，当需要处理中文，日文或者其他复合字符时，则需要用到rune类型
// byte类型是uint8的别名，对于只占用一个字节的传统ASCII编码的字符来说，完全没有问题
func main() {
	//字符类型是用单引号括起来的单个字符
	var ch byte = 'A'
	//格式化打印 使用占位符%c 输出一个byte类型的数据
	fmt.Printf("%c\n", ch)
	//rune类型等价于int32类型，可以存储带中文的复合字符
	var ch2 = '帅'
	fmt.Printf("%c\n", ch2)
}
