package main

import (
	"fmt"
	"os"
)

//练习1.1：修改echo程序输出os.Args[0]，即命令的名字。

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
