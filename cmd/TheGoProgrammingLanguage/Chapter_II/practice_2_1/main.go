package main

import (
	"GoLearn/cmd/TheGoProgrammingLanguage/Chapter_II/practice_2_1/tempconv"
	"fmt"
)

//练习 2.1： 向tempconv包添加类型、常量和函数用来处理Kelvin绝对温度的转换，
//Kelvin 绝对零度是−273.15°C，Kelvin绝对温度1K和摄氏度1°C的单位间隔是一样的。

func main() {
	fmt.Printf("AbsoluteZeroC: %v\n", tempconv.AbsoluteZeroC)
	fmt.Printf("AbsoluteZeroF: %v\n", tempconv.CToF(tempconv.AbsoluteZeroC))
	fmt.Printf("AbsoluteZeroK: %v\n", tempconv.CToK(tempconv.AbsoluteZeroC))
	fmt.Println()
	fmt.Printf("FreezingC: %v\n", tempconv.FreezingC)
	fmt.Printf("FreezingF: %v\n", tempconv.CToF(tempconv.FreezingC))
	fmt.Printf("FreezingK: %v\n", tempconv.CToK(tempconv.FreezingC))
	fmt.Println()
	fmt.Printf("BoilingC: %v\n", tempconv.BoilingC)
	fmt.Printf("BoilingF: %v\n", tempconv.CToF(tempconv.BoilingC))
	fmt.Printf("BoilingK: %v\n", tempconv.CToK(tempconv.BoilingC))
}
