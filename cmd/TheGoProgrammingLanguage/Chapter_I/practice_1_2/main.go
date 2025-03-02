package main

import (
	"fmt"
	"os"
)

//练习1.2：修改echo程序，输出参数的索引和值，每行一个。

func main() {
	for index, s := range os.Args {
		fmt.Println(index, " ", s)
	}
}
