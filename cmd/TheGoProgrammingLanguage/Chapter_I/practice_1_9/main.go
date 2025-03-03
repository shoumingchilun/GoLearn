package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
)

//练习 1.9： 修改fetch打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码。

// 命令行输入：go run main.go https://www.baidu.com
func main() {
	for _, urlStr := range os.Args[1:] {

		if hasScheme := regexp.MustCompile(`(?i)^[a-z][a-z0-9+.-]*://`).MatchString(urlStr); !hasScheme {
			urlStr = "http://" + urlStr
		}
		fmt.Println(urlStr)
		resp, err := http.Get(urlStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(resp.Status)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", urlStr, err)
			os.Exit(1)
		}
	}
}
