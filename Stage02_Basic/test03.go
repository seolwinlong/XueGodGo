package main

import (
	"fmt"
)

//将一个数组的元素的顺序进行反转。
//例如：湾前过渡小舟虚
//输出：虚舟小渡过前湾

func main() {
	arr := "湾前过渡小舟虚"
	for i, i2 := range arr {
		fmt.Println(i, string(i2))
	}
}
