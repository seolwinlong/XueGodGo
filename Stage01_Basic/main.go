package main

//包的调用：嗲用包里面的函数或变量需要遵循以下规则：
//函数或变量首字母大写表示可以被外部调用，否则只能在包内部调用
import (
	"Stage01_Basic/eat"
	test2 "Stage01_Basic/test"
	"fmt"
)

func main() {
	res := eat.Eat("面包")
	fmt.Println(res)

	//Str1首字母大写可以被调用
	s := test2.Str1
	fmt.Println(s)

	//Test函数首字母大写可以被调用
	s2 := test2.Test()
	fmt.Println(s2)

	//str2只是Test函数的返回值，不能被外部调用
}
