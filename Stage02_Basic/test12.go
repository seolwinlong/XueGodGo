package main

import "fmt"

// 匿名字段
/*所谓匿名字段（Anonymous field）是指没有名字，仅有类型的字段，也被称为嵌入字段。
通过匿名字段实现面向对象的继承关系，就是通过结构体嵌套结构体来完成的。*/

type Person struct {
	id   int
	name string
	age  int
}

type StudentA struct {
	//将结构体名作为另外一个结构体成员
	Person        //匿名字段
	name   string //同名字段
	score  float64
}

func main() {
	//在Student1对象属性中需要初始化Person信息
	//同名字段遵循就近原则
	stu := StudentA{Person{1001, "薛文龙", 37}, "大龙", 99.9}

	//匿名字段成员操作
	stu.id = 1002
	fmt.Println(stu)
	stu.score = 98
	stu.name = "文心雕龙"
	fmt.Println(stu)

	//如果想对Person中的name进行赋值可以采用以下形式
	stu.Person.name = "文心"
	fmt.Println(stu)

	//多重继承
	//多重继承指的是一个类可以继承另外一个类，而另外一个类又可以继承别的类。
	//A类继承B类，B类继承C类。
	//注意：尽量在程序中，减少多重继承，否则会增加程序的复杂度。
}
