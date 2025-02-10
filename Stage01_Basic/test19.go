package main

import "fmt"

// for循环
func main() {
	//for循环快速入门
	for i := 0; i < 10; i++ {
		fmt.Println("你好，welcome to xuegod", i)
	}

	k := 1
	for k <= 10 {
		fmt.Println("你好，welcome to xuegod", k)
		k++
	}

	//多重循环
	//1. 将一个循环放在另一个循环体内，就形成了嵌套循环。在外边的 for 称为外层循环，在里面的 for 循环称为内层循环。(建议一般使用两层，最多不要超过 3 层)
	//2. 实质上，嵌套循环就是把内层循环当成外层循环的循环体。当只有内层循环的循环条件为 false 时，才会完全跳出内层循环，才可结束外层的当次循环，开始下一次的循环。
	//3. 外层循环次数为 m 次，内层为 n 次，则内层循环体实际上需要执行 m*n 次
}
