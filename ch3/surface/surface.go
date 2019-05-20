package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	write(fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height), "data.svg")

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			write(fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy), "data.svg")
		}
	}
	write(fmt.Sprintf("</svg>"), "data.svg")
}

func corner(i, j int) (float64, float64) {
	// 求出网格单元(i,j)的顶点坐标(x, y)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 计算曲面高度z
	z := f(x, y)

	// 将x,y,z等角投射到而为SVG绘图平面上， 坐标是(sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func write(piece string, filename string) {
	data := []byte(piece)
	file, err := os.OpenFile(filename, os.O_APPEND | os.O_WRONLY, os.ModeAppend)
	defer file.Close()
	if err != nil {
		file, err = os.Create(filename)
		if err != nil {
			panic(err)
		}
		_, err = file.Write(data)
		if err != nil {
			panic(err)
		}
	} else {
		_, err = file.Write(data)
		if err != nil {
			panic(err)
		}
	}
}
