package main

import (
	"bufio"
	"fmt"
	"os"
)

//练习1.4：修改dup2程序，输出出现重复行的文件的名称。

/*提供两种理解与实现
1. 单个文件中存在重复行，则打印该文件名
2. 多个文件中存在重复行，打印存在重复行的文件
PowerShell命令为：go run main.go test1.txt test2.txt test3.txt
*/

func main() {
	files := os.Args[1:]
	//判别单个文件内行重复
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		findDupLinesInSingleFile(f)
		f.Close()
	}

	//分隔线
	fmt.Println("*********************************")

	//判别多个文件内是否存在行重复
	flags := make(map[string]map[string]bool)
	//map:<行内容<出现在的文件名，该文件内部是否重复>>
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		findDupLinesAcrossFiles(f, flags)
		f.Close()
	}
	result := make(map[string]map[string]struct{})
	//记录结果：map:<文件名:行>
	for line, fileNames := range flags {
		if len(fileNames) > 1 {
			//说明同时存在两个文件中
			for fileName := range fileNames {
				if lines, exist := result[fileName]; exist {
					lines[line] = struct{}{}
				} else {
					result[fileName] = map[string]struct{}{line: {}}
				}
			}
		}
		for fileName, isDup := range fileNames {
			if isDup {
				//说明在该文件中已重复
				if lines, exist := result[fileName]; exist {
					lines[line] = struct{}{}
				} else {
					result[fileName] = map[string]struct{}{line: {}}
				}
			}
		}
	}
	//输出结果
	for fileName, lines := range result {
		fmt.Println(fileName)
		for line := range lines {
			fmt.Println("\t", line)
		}
	}
}

func findDupLinesAcrossFiles(f *os.File, flags map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		str := input.Text()
		if files, exist := flags[str]; exist {
			//该行已出现过
			if _, exist2 := files[f.Name()]; exist2 {
				//该行在当前文件中已经出现过
				files[f.Name()] = true
			} else {
				//该行在当前文件中未出现过
				files[f.Name()] = false
			}
		} else {
			flags[str] = map[string]bool{f.Name(): false}
		}
	}
}

func findDupLinesInSingleFile(f *os.File) {
	flags := make(map[string]bool)
	//flags记录出现过的行，nil说明未出现过，false说明仅出现过一次，bool说明重复
	flag := false
	//flag标记该文件中是否存在重复行
	input := bufio.NewScanner(f)
	for input.Scan() {
		str := input.Text()
		if _, exist := flags[str]; exist {
			//已出现过
			flags[str] = true
			flag = true
		} else {
			//首次出现
			flags[str] = false
		}
	}
	if flag {
		fmt.Println(f.Name())
		for k, b := range flags {
			if b {
				fmt.Println("\t", k)
			}
		}
	}
}
