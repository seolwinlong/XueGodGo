package model

import "fmt"

type person struct {
	Name string
	age  int //其他包不能直接访问
	sal  float64
}

// 写一个工厂模式函数，相当于构造函数
func newPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// 为了访问age和sal我们编写一对SetXxx和GetXxx的方法
func (p *person) setAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确")
		//给程序员一个默认值
	}
}

func (p *person) getAge() int {
	return p.age
}

func (p *person) setSal(sal float64) {
	if sal >= 3000 && sal <= 30000 {
		p.sal = sal
	} else {
		fmt.Println("薪水范围不正确..")
	}
}

func (p *person) getSal() float64 {
	return p.sal
}
