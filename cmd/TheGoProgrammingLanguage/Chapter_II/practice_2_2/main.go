package main

import (
	"GoLearn/cmd/TheGoProgrammingLanguage/Chapter_II/practice_2_2/weightconv"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//练习 2.2： 写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数，如果缺省的话则是从标准输入读取参数，
//然后做类似Celsius和Fahrenheit的单位转换，长度单位可以对应英尺和米，重量单位可以对应磅和公斤等。

// 命令行参数：go run main.go 12.23 54 322
func main() {
	if args := os.Args[1:]; len(args) != 0 {
		for _, arg := range args {
			if num, err := strconv.ParseFloat(arg, 64); err == nil {
				fmt.Println(weightconv.Pound(num), "=", weightconv.Pound(num).ToKG())
				fmt.Println(weightconv.Kilogram(num), "=", weightconv.Kilogram(num).ToLB())
			} else {
				fmt.Println(fmt.Errorf("入参格式异常:%v", arg))
			}
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if num, err := strconv.ParseFloat(line, 64); err == nil {
				fmt.Println(weightconv.Pound(num), "=", weightconv.Pound(num).ToKG())
				fmt.Println(weightconv.Kilogram(num), "=", weightconv.Kilogram(num).ToLB())
			} else {
				fmt.Println(fmt.Errorf("入参格式异常:%v", line))
			}
		}
	}
}
