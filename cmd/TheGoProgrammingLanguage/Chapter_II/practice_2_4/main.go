package main

import (
	"GoLearn/cmd/TheGoProgrammingLanguage/Chapter_II/practice_2_4/popcount"
	"fmt"
	"math/rand"
)

//练习 2.4： 用移位算法重写PopCount函数，每次测试最右边的1bit，然后统计总数。比较和查表算法的性能差异。

// 在终端进入popcount文件夹，命令行执行go test --bench=.
func main() {
	for i := 0; i < 1000; i++ {
		num := uint64(rand.Int63())
		fmt.Println(num)
		fmt.Println(popcount.PopCount(num))
		fmt.Println(popcount.PopCount2(num))
	}
}
