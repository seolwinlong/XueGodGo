package main

import "fmt"

// BubbleSort 切片实现冒泡排序
func BubbleSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}
func main() {
	slice := []int{9, 1, 5, 6, 10, 8, 3, 7, 2, 4}
	BubbleSort(slice)
	//主调函数中实参的值被修改
	fmt.Println(slice)
}

//总结：建议在开发中使用切片代替数组。
