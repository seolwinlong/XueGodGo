package main

import (
	"fmt"
	"time"
)

//并发编程
//所谓并发编程是指在一台处理器上“同时”处理多个任务。
/*并发：一个CPU上能同时执行多项任务，在很短时间内，CPU来回切换任务执行(在某段很短时间内执行程序a，然后又迅速得切换到程序b去执行)，
有时间上的重叠（宏观上是同时的，微观仍是顺序执行）,这样看起来多个任务像是同时执行，这就是并发。*/

/*
并行：指在同一时刻只能有一条指令执行，但多个进程指令被快速地轮换执行，使得在宏观上具有多个进程同
时执行的效果，但在微观上并不是同时执行的，只是把时间分成若干段，通过CPU时间片轮转使多个进程快速交替的执行。
*/

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("newroutine:i=%d\n", i)
		time.Sleep(time.Second) //延时一秒
	}
}

func main() {
	//创建一个goroutine，启动另外一个任务
	go newTask()
	i := 0
	for {
		i++
		fmt.Printf("mainroutine:i=%d\n", i)
		time.Sleep(time.Second) //延时1s
	}
}
