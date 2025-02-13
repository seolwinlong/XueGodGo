package main

import (
	"fmt"
)

// 案例1：一个班级有5个学生，要求分别输入成绩，求最好成绩，平均值
func main() {
	var arr [5]int
	for i := 0; i <= 4; i++ {
		fmt.Printf("请输入第%d个学生成绩：", i+1)
		_, err := fmt.Scanf("%d", &arr[i])
		if err != nil {
			return
		}
	}
	bestScore := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > bestScore {
			bestScore = arr[i]
		}
	}
	fmt.Printf("最好成绩是：%d\n", bestScore)

	var totalScore int

	for i := 0; i < len(arr); i++ {
		totalScore += arr[i]
		fmt.Println(i)
	}
	avgScore := totalScore / len(arr)
	fmt.Printf("平均成绩为：%d", avgScore)
}
