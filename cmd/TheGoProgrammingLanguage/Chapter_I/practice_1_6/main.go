package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// 练习 1.6： 修改Lissajous程序，修改其调色板来生成更丰富的颜色，
// 然后修改SetColorIndex的第三个参数，看看显示结果吧。
var palette = []color.Color{
	color.Black,
	color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff},
	color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff},
	color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff},
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	//直接创建文件，powershell下使用>输出可能存在异常导致gif无法被正确解析
	f, err := os.Create("out.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lissajous(f)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}

	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		//palette中的第一个颜色会变成默认的背景颜色
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Intn(3)+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
