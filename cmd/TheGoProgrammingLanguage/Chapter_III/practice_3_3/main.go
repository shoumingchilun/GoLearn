// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
	"os"
)

//练习 3.3： 根据高度给每个多边形上色，那样峰值部将是红色（#ff0000），谷部将是蓝色（#0000ff）。

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

const RGBHeightSize = 0.5 //R使用ff时对应的高度

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
			ax, ay, aheight := corner(i+1, j)
			bx, by, bheight := corner(i, j)
			cx, cy, cheight := corner(i, j+1)
			dx, dy, dheight := corner(i+1, j+1)

			if !isValid(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}

			RGBheight := (aheight + bheight + cheight + dheight) / 4
			if RGBheight > RGBHeightSize {
				RGBheight = RGBHeightSize
			}
			if RGBheight < -RGBHeightSize {
				RGBheight = -RGBHeightSize
			}
			R := int(RGBheight*0xff/(2*RGBHeightSize) + 0xff/2)
			B := int(RGBheight*-0xff/(2*RGBHeightSize) + 0xff/2)
			_, err := file.WriteString(fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: #%02x00%02x'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, R, B))
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

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func isValid(nums ...float64) bool {
	for _, num := range nums {
		if math.IsNaN(num) || math.IsInf(num, 0) {
			return false
		}
	}
	return true
}
