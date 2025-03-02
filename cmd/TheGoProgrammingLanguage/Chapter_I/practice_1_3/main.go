package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//练习1.3：尝试测量可能低效的程序和使用strings.Join的程序在执行时间上的差异。（1.6节有time包，11.4节展示如何撰写系统性的性能评估测试。）

func main() {
	//构造用例
	nums := make([]string, 100000)
	for i := 0; i < len(nums); i++ {
		nums[i] = strconv.Itoa(rand.Int())
	}
	//使用+
	begin1 := time.Now().Unix()
	useStringPlus(nums[:])
	end1 := time.Now().Unix()
	//使用join
	begin2 := time.Now().Unix()
	useStringsJoin(nums[:])
	end2 := time.Now().Unix()
	//输出结果
	fmt.Println("第一次耗时：", end1-begin1, "s")
	fmt.Println("第二次耗时：", end2-begin2, "s")
}

func useStringsJoin(nums []string) string {
	sum2 := strings.Join(nums[:], "")
	return sum2
}

func useStringPlus(nums []string) string {
	sum1 := ""
	for _, num := range nums {
		sum1 += num
	}
	return sum1
}
