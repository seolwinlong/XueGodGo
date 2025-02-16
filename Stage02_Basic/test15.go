package main

import "fmt"

// PersonC 方法继承
// 如果一个结构体通过匿名字段作为另外一个结构体成员，该结构体的方法也可以被另外继承关系的结构体使用，称为方法继承。
type PersonC struct {
	id   int
	name string
	age  int
}

type StudentC struct {
	PersonC
	score int
}

func (p *PersonC) PrintInfo() {
	fmt.Printf("name is %s, age is %d\n", p.name, p.age)
}

func main() {
	stu := StudentC{PersonC{1001, "winlong", 37}, 100}
	//就近原则：先找本作用域的方法，找不到再用继承的方法
	stu.PrintInfo()
	//显式调用继承的方法
	stu.PersonC.PrintInfo()
}
