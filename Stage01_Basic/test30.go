package main

import (
	"errors"
	"fmt"
)

// 错误处理
/*异常是当Go语言在执行时检测到一个错误时，程序就无法继续执行了，出现错误的提示，这就是所谓
的"异常"。
1. 在默认情况下，当发生错误后(panic) ,程序就会退出（崩溃.）。 panic惊恐
2. 如果我们希望：当发生错误后，可以捕获到错误，并进行处理，保证程序可以继续执行。还可以在
捕获到错误后，给管理员一个提示(邮件,短信。。。）*/
//Go 中引入的处理方式为：defer, panic, recover
//这几个异常的使用场景可以这么简单描述：Go 中可以抛出一个 panic 的异常，然后在 defer 中通过recover 捕获这个异常，然后正常处理。
//使用error处理错误信息

func ErrorTest(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("除数不能为0")
		return
	}
	result = a / b
	return
}
func main() {
	//使用result接收结果
	//使用err接收错误信息
	result, err := ErrorTest(1, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
