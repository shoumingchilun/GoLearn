package main

import (
	"GoLearn/cmd/TheGoProgrammingLanguage/Chapter_II/practice_2_5/popcount"
	"fmt"
	"math/rand"
)

//练习 2.5： 表达式x&(x-1)用于将x的最低的一个非零的bit位清零。使用这个算法重写PopCount函数，然后比较性能。

// 在终端进入popcount文件夹，命令行执行go test --bench=.
func main() {
	for i := 0; i < 1000; i++ {
		num := uint64(rand.Int63())
		fmt.Println(num)
		fmt.Println(popcount.PopCount(num))
		fmt.Println(popcount.PopCount2(num))
	}
}
