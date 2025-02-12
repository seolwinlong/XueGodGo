package main

import "fmt"

// 数组初始化示例
func main() {
	//1、指定长度的初始化
	var arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)
	//或者
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	//2、如果数组长度不确定，可以使用...代替数组的长度，编译器会根据元素个数自行推断数组的长度
	var arr3 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)
	//或者
	arr4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr4)

	//3、不指定长度，key：value方式初始化
	arr5 := [...]int{1: 100, 2: 200, 3: 300, 4: 100}
	fmt.Println(arr5)

	//4、创建数组不进行初始化，通过数组索引单独进行赋值
	var arr6 [3]int
	arr6[0] = 1000
	arr6[2] = 2000
	fmt.Println(arr6)

	//数组遍历forr
	for index, value := range arr5 {
		fmt.Println(index, value)
	}

	//扮演奇怪的神仙
	PlayStrangeGods := [...]string{"孙悟空", "猪八戒", "沙和尚"}

	for i, god := range PlayStrangeGods {
		fmt.Printf("i=%v v=%v\n", i, god)
		fmt.Printf("Play_strange_gods[%d]=%v\n", i, PlayStrangeGods[i])
	}

	//不需要index可以使用_下划线忽略返回值
	for _, god := range PlayStrangeGods {
		fmt.Printf("元素的值是:%v\n", god)
	}
}
