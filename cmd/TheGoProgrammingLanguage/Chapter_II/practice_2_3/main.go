package main

import (
	"GoLearn/cmd/TheGoProgrammingLanguage/Chapter_II/practice_2_3/popcount"
	"fmt"
	"math/rand"
)

//练习 2.3： 重写PopCount函数，用一个循环代替单一的表达式。比较两个版本的性能。
//（11.4节将展示如何系统地比较两个不同实现的性能。）

// 在终端进入popcount文件夹，命令行执行go test --bench=.
func main() {
	for i := 0; i < 1000; i++ {
		num := uint64(rand.Int63())
		fmt.Println(num)
		fmt.Println(popcount.PopCount(num))
		fmt.Println(popcount.PopCount2(num))
	}
}
