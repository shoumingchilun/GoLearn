// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
	"os"
)

//练习 3.2： 试验math包中其他函数的渲染图形。你是否能输出一个egg box、moguls或a saddle图案?

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	file, err := os.Create("temp.svg")
	if err != nil {
		return
	}
	_, err = file.WriteString(fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height))
	if err != nil {
		return
	}

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			if !isValid(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}

			_, err := file.WriteString(fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy))
			if err != nil {
				return
			}
		}
	}
	_, err = file.WriteString(fmt.Sprint("</svg>"))
	if err != nil {
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(file)
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	return aSaddle(x, y)
	//return mogulHeight(x, y)
	//return eggBox(x, y)
}

func eggBox(x, y float64) float64 {
	return math.Sin(x) * math.Sin(y) / 4
}

// Mogul 定义雪包参数结构体
type Mogul struct {
	X, Y  float64 // 中心坐标
	A     float64 // 振幅（高度）
	Sigma float64 // 标准差（宽度）
}

// 定义雪包参数组
var moguls = []Mogul{
	{X: -10, Y: -10, A: 1.0, Sigma: 3.0},
	{X: 10, Y: 10, A: 1.0, Sigma: 3.0},
	{X: -10, Y: 10, A: 0.75, Sigma: 2.5},
	{X: 10, Y: -10, A: 0.75, Sigma: 2.5},
}

func mogulHeight(x, y float64) float64 {
	var z float64
	for _, m := range moguls {
		dx := x - m.X
		dy := y - m.Y
		// 高斯函数: A * exp(-(dx² + dy²)/(2σ²))
		z += m.A * math.Exp(-(dx*dx+dy*dy)/(2*m.Sigma*m.Sigma))
	}
	return z
}

func aSaddle(x, y float64) float64 {
	return (x*x - y*y) / 300
}

func isValid(nums ...float64) bool {
	for _, num := range nums {
		if math.IsNaN(num) || math.IsInf(num, 0) {
			return false
		}
	}
	return true
}
