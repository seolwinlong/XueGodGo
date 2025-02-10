package main

import "fmt"

//goto无条件跳转
/*1. Go 语言的 goto 语句可以无条件地转移到程序中指定的行。
  2. goto 语句通常与条件语句配合使用。可用来实现条件转移，跳出循环体等功能。
  3. 在 Go 程序设计中一般不主张使用 goto 语句， 以免造成程序流程的混乱，使理解和调试程序都产生困难。*/
func main() {
	n := 30
	fmt.Println("ok1")
	if n > 20 {
		goto label1
	}
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
label1:
	fmt.Println("ok5")
	fmt.Println("ok6")
	fmt.Println("ok7")
}
