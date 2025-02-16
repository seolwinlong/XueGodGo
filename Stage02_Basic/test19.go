package main

import "fmt"

// Humaner 为了保护你的Go语言职业生涯，请牢记接口（interface）是一种类型。
type Humaner interface {
	// SayHello SayHello只有定义没有实现
	SayHello()
}

type StudentE struct {
	name string
	age  int
}

func (stu StudentE) SayHello() {
	fmt.Printf("%s is %d years old\n", stu.name, stu.age)
}

type Teacher struct {
	name    string
	subject string
}

func (t Teacher) SayHello() {
	fmt.Printf("%s teaching %s \n", t.name, t.subject)
}

func main() {
	stu := StudentE{"Tom", 20}
	tea := Teacher{"David", "Math"}
	fmt.Println(stu, tea)

	//定义接口类型变量
	var h Humaner

	//如果对象的方法为值接收者，直接将对象赋值给接口变量
	h = stu
	h.SayHello()

	//如果对象的方法为指针接收者，将对象取地址后赋值给接口变量
	h = &tea
	h.SayHello()
}
