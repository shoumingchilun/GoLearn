package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

//练习 1.8： 修改fetch这个范例，如果输入的url参数没有 http:// 前缀的话，为这个url加上该前缀。你可能会用到strings.HasPrefix这个函数。

// 命令行输入：go run main.go www.baidu.com
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
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", urlStr, err)
			os.Exit(1)
		}
	}
}
