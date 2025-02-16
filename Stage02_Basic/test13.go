package main

import "fmt"

type People struct {
	id   int
	name string
	age  int
}

type Winlong struct {
	People
	score int
}

// Play 值对象接收者
// 在指针对象方法接收者修改不影响源对象的信息
func (w Winlong) Play() string {
	w.name = "小薛"
	return "正在打英雄联盟手游！"
}

func main() {
	w := Winlong{People{1001, "薛文龙", 37}, 100}
	fmt.Println(w)
	fmt.Println(w.Play())
	fmt.Println(w)
}
