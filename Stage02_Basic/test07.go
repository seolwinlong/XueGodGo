package main

import "fmt"

func Demo(slice []int) {
	//修改切片元素的值
	slice[0] = 123
	fmt.Println(slice)
	fmt.Printf("%p\n", slice)
}

//切片的遍历fori和forr

func main() {
	//使用fori
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Println("------------------")

	//使用forr
	for _, value := range arr {
		fmt.Println(value)
	}

	//切片的扩容
	slice := []int{1, 2, 3, 4, 5}
	newSlice := append(slice, 6, 6, 6)
	fmt.Println(newSlice)

	//切片的截取
	fmt.Println(slice[0])
	s1 := slice[:]
	fmt.Println(len(s1), cap(s1))

	slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//设置截取的起始索引
	s2 := slice2[3:]
	fmt.Println(s2)
	fmt.Println("长度：", len(s2), "容量：", cap(s2))

	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//设置截取的结束索引 左闭右开
	s3 := slice3[:5]
	fmt.Println(s3)
	fmt.Println("长度：", len(s3), "容量：", cap(s3))

	slice4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//设置截取的起始索引和结束索引 左闭右开
	s4 := slice4[3:5]
	fmt.Println(s4)
	fmt.Println("长度：", len(s4), "容量：", cap(s4))

	slice5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//设置截取的起始索引和结束索引并指定容量
	//容量：cap=max-low
	s5 := slice5[3:5:5]
	fmt.Println(s5)
	fmt.Println("长度：", len(s5), "容量：", cap(s5))

	slice6 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//创建一个长度为5的元素大小的切片
	s := make([]int, 5)
	copy(s, slice6)
	fmt.Println(s)

	//复制后的切片和源切片为不同切片，即内存地址不同。修改一个不会影响另外一个。
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//创建一个长度为10的元素大小的切片
	dst := make([]int, 10)
	//切片拷贝
	copy(dst, src)
	//打印切片地址
	fmt.Printf("源切片：%p\n", src)
	fmt.Printf("目标切片：%p\n", dst)

	//切片也可以作为函数参数 ，为值传递（go语言所有的参数都是值传递）。
	//形参可以间接修改实参的值。
	slice7 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Demo(slice7)
}
