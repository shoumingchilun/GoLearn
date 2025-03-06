package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//练习 1.12： 修改Lissajour服务，从URL读取变量，比如你可以访问 http://localhost:8000/?cycles=20 这个URL，
//这样访问可以将程序里的cycles默认的5修改为20。字符串转换为数字可以调用strconv.Atoi函数。你可以在godoc里查看strconv.Atoi的详细说明。

// 浏览器访问：http://localhost:8000/?size=400&delay=5&nframes=250&res=0.005&cycles=20
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()

	cycles, done := getQueryParameter(writer, query, "cycles", 5, parseInt)
	if done {
		return
	}
	res, done := getQueryParameter(writer, query, "res", 0.001, parseFloat)
	if done {
		return
	}
	size, done := getQueryParameter(writer, query, "size", 100, parseInt)
	if done {
		return
	}
	nframes, done := getQueryParameter(writer, query, "nframes,", 64, parseInt)
	if done {
		return
	}
	delay, done := getQueryParameter(writer, query, "delay", 8, parseInt)
	if done {
		return
	}

	lissajous(writer, cycles, res, size, nframes, delay)
}

// 类型转换函数
func parseInt(value string) (int, error) {
	return strconv.Atoi(value)
}

func parseFloat(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

// 泛型方法：从查询参数中获取指定类型的值
func getQueryParameter[T any](writer http.ResponseWriter, query url.Values, key string, defaultValue T, parser func(string) (T, error)) (T, bool) {
	if value := query.Get(key); value != "" {
		if parsedValue, err := parser(value); err == nil {
			return parsedValue, false
		} else {
			http.Error(writer, fmt.Sprintf("Invalid %s value", key), http.StatusBadRequest)
			return defaultValue, true
		}
	}
	return defaultValue, false
}

var palette = []color.Color{
	color.Black,
	color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff},
	color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff},
	color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff},
}

func lissajous(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {
	//const (
	//	cycles  = 5     // number of complete x oscillator revolutions
	//	res     = 0.001 // angular resolution
	//	size    = 100   // image canvas covers [-size..+size]
	//	nframes = 64    // number of animation frames
	//	delay   = 8     // delay between frames in 10ms units
	//)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}

	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		//palette中的第一个颜色会变成默认的背景颜色
		for t := 0.0; t < float64(cycles*2)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), uint8(rand.Intn(3)+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
