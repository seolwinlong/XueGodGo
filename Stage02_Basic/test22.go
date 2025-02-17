package main

/*interface的变量里面可以存储任意类型的数值(该类型实现了interface)。那么怎么反向知道这个变量里面实际保存了的是哪个类型的对象呢？
目前常用的有两种方法：
comma-ok断言
switch测试
Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false。*/

/*func main() {
//定义一个空接口切片
var slice interface{}
i := append(slice, 3.14, true, [3]int{1, 2, 3}, "hello", 10, 123.456, "你瞅啥")
fmt.Println(i)
//通过类型断言 获取数据类型
//空接口数据.(数据类型)
value, ok := slice[2].(int)
if ok {
	fmt.Println(value)
} else {
	fmt.Println("不是整型数据")
}*/

//注意：
//类型断言就是将接口类型的值(x)，转换成类型(T)。格式为：x.(T);
//类型断言的必要条件就是x是接口类型，非接口类型的x不能做类型断言；
//T可以是非接口类型，如想断言合法，则T必须实现x的接口；
//也可以是接口，则x的动态类型也应该实现接口T；
//类型断言如果非法，运行失败时会导致错误，请谨慎使用。
//}
