package main

import (
	"fmt"
	"time"
)

// Channel
/*channel是Go语言中的一个核心类型，可以把它看成管道。并发核心单元通过其就可以发送或者接收数据进行通讯，在一定程度上又进一步降低编程的难度。
channel是一个数据类型，主要用来解决Go程的同步问题以及Go程之间数据共享（数据传递）的问题。
goroutine运行在相同的地址空间，因此访问共享内存必须做好同步。goroutine 奉行通过通信来共享内存，而不是共享内存来通信。
引⽤类型 channel可用于多个 goroutine 通讯。其内部实现了同步，确保并发安全。*/

func main() {
	//创建channel
	//channel也是一个对应make创建的底层数据结构的引用。
	//chan是创建channel所需使用的关键字。Type 代表指定channel收发数据的类型。
	/*make(chan Type) //等价于make(chan Type,0) 无缓冲的通道
	make(chan Type,capacity) //等价于make(chan Type,容量) 有缓冲的通道*/
	//无缓冲的通道（unbuffered channel）是指在接收前没有能力保存任何数据值的通道。
	//阻塞：由于某种原因数据没有到达，当前go程（线程）持续处于等待状态，直到条件满足，才解除阻塞。
	//同步：在两个或多个go程（线程）间，保持数据内容一致性的机制。

	c := make(chan int) //创建无缓冲通道c
	//内置函数len返回未被读取的缓冲元素数量，cap返回缓冲区的大小
	fmt.Printf("len(c)=%d,cap(c)=%d\n", len(c), cap(c))
	go func() {
		defer fmt.Println("子go程结束")
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Printf("子go程正在运行[%d]:len(c)=%d,cap(c)=%d\n", i, len(c), cap(c))
		}
	}()
	time.Sleep(2 * time.Second) //延时两秒
	for i := 0; i < 3; i++ {
		num := <-c //从c中接收数据，并赋值给num
		fmt.Println("num=", num)
	}
	fmt.Println("main进程结束")
}
