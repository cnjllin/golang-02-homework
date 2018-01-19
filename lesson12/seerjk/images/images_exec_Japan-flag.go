package main

import (
	"golang.org/x/tour/pic"
	"image/color"
	"image"
	"math"
)

type Image struct {
	Width, Height int
}

type Point struct {
	X, Y int
}

func (p *Point) PointDistance(pOther Point) float64 {
	return math.Sqrt(math.Pow(float64(p.X-pOther.X), 2.0) + math.Pow(float64(p.Y-pOther.Y), 2.0))
}

type Circle struct {
	p Point  // 圆心
	R int 	 // 半径
}

/*
type Image interface {
	// ColorModel returns the Image's color model.
	ColorModel() color.Model
	// Bounds returns the domain for which At can return non-zero color.
	// The bounds do not necessarily contain the point (0, 0).
	Bounds() Rectangle
	// At returns the color of the pixel at (x, y).
	// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
	// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
	At(x, y int) color.Color
}
*/

func (im *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.Width, im.Height)
}

func (im *Image) At(x, y int) color.Color {
	// 绘制日本国旗
	circlePoint := Point{
		X: im.Width / 2,
		Y: im.Height / 2,
	}

	currentPoint := Point{
		X: x,
		Y: y,
	}

	c := Circle{
		p: circlePoint,
		R: int(math.Min(float64(im.Width), float64(im.Height)) / 5),
	}
	// 判断方法 点到圆心的距离 <= 半径
	if currentPoint.PointDistance(circlePoint) <= float64(c.R) {
		return color.RGBA{255, 0, 0, 255}
	} else {
		return color.RGBA{255, 255, 255, 255}
	}
}

func main() {
	m := Image{200, 200}
	pic.ShowImage(&m)
}
