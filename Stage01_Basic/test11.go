package main

import "fmt"

// 字符串类型，用双引号括起来的字符是字符串类型
// 在Go中的字符串，都是用UTF-8字符集编码
func main() {
	var str = "刘源真帅"
	fmt.Printf("str=%s\n", str)

	//字符串拼接
	str1 := "hello"
	str2 := "学神"
	//len(字符串)计算字符串中字符的个数
	fmt.Printf("str1的长度是：%d\n", len(str1))
	fmt.Printf("str2的长度是：%d\n", len(str2))
	str3 := str2 + str1
	fmt.Println(str3)
}
