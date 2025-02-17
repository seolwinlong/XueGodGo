package main

import (
	"fmt"
	"os"
)

// 文件处理

// 根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666的文件，返回的文件对象是可读写的。
// func Create(name string) (file *file,err Error)

// 根据文件描述符创建相应的文件，返回一个文件对象
// func NewFile(fd uintptr,name string) *file

//该方法打开一个名称为name的文件，但是是只读方式，内部实现其实调用了OpenFile。
//func Open(name string) (file *File, err Error)

//打开名称为name的文件，flag是打开的方式，只读、读写等，perm是权限
//func OpenFile(name string, flag int, perm uint32) (file *File, err Error)

func main() {
	//创建文件
	file, err := os.Create("C:/Users/winlo/desktop/test.txt")
	if err != nil {
		fmt.Println("Create err", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Close err", err)
		}
	}(file)
}
