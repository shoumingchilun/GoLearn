package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//练习 1.10： 找一个数据量比较大的网站，用本小节中的程序调研网站的缓存策略，对每个URL执行两遍请求，
//查看两次时间是否有较大的差别，并且每次获取到的响应内容是否一致，修改本节中的程序，将响应结果输出到文件，以便于进行对比。

//命令行参数： go run main.go https://chat.openai.com/

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
		go fetch(url, ch)
	}
	for i := 0; i < len(os.Args[1:])*2; i++ {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	var nbytes int64
	fileName := strings.ReplaceAll(url+"_"+strconv.Itoa(int(time.Now().UnixNano()))+".txt", "://", "_")
	fileName = strings.ReplaceAll(fileName, "/", "_")
	f, err := os.Create(fileName)
	if err == nil {
		nbytes, err = io.Copy(f, resp.Body)
		defer f.Close()
		defer resp.Body.Close()
	}
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
