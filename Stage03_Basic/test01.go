package main

import (
	"fmt"
	"strings"
)

//字符串处理函数

func main() {
	//字符串包含子串
	fmt.Println(strings.Contains("xueshenit", "it"))
	fmt.Println(strings.Contains("xueshenit", "winlong"))
	fmt.Println(strings.Contains("xueshenit", ""))
	fmt.Println(strings.Contains("", "it"))

	//字符串查找位置
	fmt.Println(strings.Index("xueshenit", "it"))
	fmt.Println(strings.Index("xueshenit", "winlong"))
	fmt.Println(strings.Index("xueshenit", ""))

	//字符串拼接
	s := []string{"xue", "shen", "it"}
	fmt.Println(strings.Join(s, ","))

	//字符串去首尾字符
	//去掉首尾的！
	fmt.Printf("[%q]\n", strings.Trim(" !!! xueshenit!xueshenit! !!!", "! "))
	//去掉首尾的空格
	fmt.Printf("[%q]\n", strings.Trim(" hello ", " "))
}
