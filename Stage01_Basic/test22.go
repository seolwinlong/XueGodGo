package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 函数基本用法
func main() {
	//使用时间戳作为随机数种子进行混淆
	//golang内置函数
	rand.NewSource(time.Now().UnixNano())
	//限定100以内的随机数
	value := rand.Intn(100)
	fmt.Println(value)
}
