package main

import "fmt"

// 算术运算符
func main() {
	a := 10
	b := 20
	fmt.Println("a+b的值为：", a+b)
	fmt.Println("a-b的值为：", a-b)
	fmt.Println("a*b的值为：", a*b)
	fmt.Println("a/b的值为：", a/b)
	fmt.Println("a取模b的值为：", a%b)
	a++
	fmt.Printf("a++的值为：%v\n", a)
	//为了方便测试，a这里重新赋值为10
	a = 10
	a--
	fmt.Printf("a--的值为：%v\n", a)
	//自增自减运算符只能作为独立语句，不能用于表达式。
	//++和--只能放在变量的后面使用，不支持放在前面
}
