package main

import "fmt"

// Go语言字典和数组切片一样，是来保存一组相同的数据类型的。
// Go语言字典可以通过key键来获取value值，map为映射关系容器，采用散列（hash）实现。

func test(m map[int]string) {
	//修改map中的value
	m[1001] = "伊泽瑞尔"
	//删除map的数据
	delete(m, 1002)
	//添加数据
	m[1008] = "亚索"
	fmt.Printf("%p\n", m)
	//fmt.Println(m)
}

func main() {
	//map在定义时不允许重复
	var m = map[int]string{1001: "刘能", 1002: "赵四", 1008: "谢广坤"}
	fmt.Println(m)

	//map遍历
	//map存储时无序的，无法决定返回顺序，所以每次打印结果的顺序有可能不同
	for key, value := range m {
		fmt.Println(key, value)
	}

	m[1020] = "王老七"    //添加
	m[1002] = "亚洲舞王赵四" //修改
	//根据键打印值
	fmt.Println(m[1001])
	fmt.Println(m[1002])
	fmt.Println(m[1008])

	//值，存在判断 := 字典[键]
	//如果map中的key存在 返回值 value表示具体key对应的值 okbool类型 表示表示key存在条件
	value, ok := m[1002]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("key不存在")
	}

	fmt.Println("删除前：")
	fmt.Println(m)
	//删除操作
	delete(m, 1008)
	fmt.Println("删除后：")
	fmt.Println(m)

	//字典也可以作为函数参数 ，作为参数为值传递。
	//形参可以间接修改实参的值。
	k := make(map[int]string)
	k[1001] = "劫"
	k[1002] = "盖伦"
	k[1005] = "薇恩"

	test(m)
	fmt.Printf("%p\n", k)
	fmt.Println(k)
}
