package main

import "fmt"

// Usb 多态
// 在 Go 语言，多态特征是通过接口实现的。可以按照
// 统一的接口来调用不同的实现。这时接口变量就呈现不同的形态。
// 声明/定义一个接口
type Usb interface {
	// Start 声明两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	name string
}

// Start 让Phone实现Usb接口方法
func (p *Phone) Start() {
	fmt.Println("手机开始工作")
}

func (p *Phone) Stop() {
	fmt.Println("手机停止工作")
}

type Camera struct {
	name string
}

// Start 让Camera实现Usb接口的方法
func (c *Camera) Start() {
	fmt.Println("相机开始工作")
}

func (c *Camera) Stop() {
	fmt.Println("相机停止工作")
}

func main() {

	phone := Phone{"苹果"}
	phone.Start()
	phone.Stop()

	camera := Camera{"科达"}
	camera.Start()
	camera.Stop()
}
