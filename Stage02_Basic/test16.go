package main

import "fmt"

//方法重写
//如果子类(结构体)中的方法名与父类(结构体)中的方法名同名，在调用的时候是先调用子类(结构体)中的
//方法，这就方法的重写。所谓的重写：就是子类(结构体)中的方法，将父类中的相同名称的方法的功能重新给改写。

type PersonD struct {
	id   int
	name string
	age  int
}

type StudentD struct {
	PersonD
	score int
}

func (p PersonD) printInfo() {
	fmt.Printf("Person->id:%d name:%s age:%d \n", p.id, p.name, p.age)
}

func (s StudentD) printInfo() {
	fmt.Printf("Student->id:%d name:%s age:%d \n", s.id, s.name, s.age)
}

func main() {
	student := StudentD{PersonD{1, "student", 20}, 30}
	//就近原则：先找本作用域的方法，找不到再用继承的方法
	student.printInfo()
	//显式调用继承的方法
	student.PersonD.printInfo()
}
