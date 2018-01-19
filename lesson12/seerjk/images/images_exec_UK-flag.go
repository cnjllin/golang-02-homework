package main

import (
	"golang.org/x/tour/pic"
	"image/color"
	"image"
	"math"
)

type Point struct {
	X, Y float64
}

type Line struct {
	PA, PB Point
}

func (l *Line) DistancToPoint(p Point) float64 {
	// 直线的两点(x1,y1)(x2,y2)式 (X - x1)/(x2 - x1) = (Y - y1)/(y2 - y1)
	// 直线方程 Ax + By +C = 0
	// A = 1/(x2-x1)
	// B = 1/(y1-y2)
	// C = y1/(y2-y1) - x1/(x2-x1)

	A := 1.0 / (l.PB.X - l.PA.X)
	B := 1.0 / (l.PA.Y - l.PB.Y)
	C := l.PA.Y/(l.PB.Y-l.PA.Y) - l.PA.X/(l.PB.X-l.PA.X)

	// 点到直线(直线方程 Ax + By +C = 0)距离公式 abs(A*x0 + B*y0 + C) / (sqrt(A^2+B^2))
	return math.Abs(A*p.X+B*p.Y+C) / math.Sqrt(math.Pow(A, 2)+math.Pow(B, 2))
}

type Image struct {
	Width, Height int
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
	// 绘制红色米字
	currentPoint := Point{float64(x), float64(y)}
	// 米字中间的十字
	halfX := im.Width / 2
	halfY := im.Height / 2

	// 米字线的 1/2 的宽度
	crossWidth := int(math.Min(float64(im.Width), float64(im.Height)) / 20)

	// 米字的两个斜线
	lineA := Line{Point{0,0}, Point{float64(im.Width), float64(im.Height)}}
	lineB := Line{Point{0, float64(im.Height)}, Point{float64(im.Width), 0}}

	//fmt.Println(lineA.DistancToPoint(currentPoint))

	if (x >= halfX-crossWidth && x <= halfX+crossWidth) ||
		y >= halfY-crossWidth && y <= halfY+crossWidth {
		//	绘制十字
		return color.RGBA{255, 0, 0, 255}
	} else if lineA.DistancToPoint(currentPoint) <= float64(crossWidth) ||
		lineB.DistancToPoint(currentPoint) <= float64(crossWidth) {
		//	绘制米字的斜线
		return color.RGBA{255, 0, 0, 255}
	} else {
		return color.RGBA{4,40,123, 255}
	}
}

func main() {
	m := Image{300, 200}
	pic.ShowImage(&m)
}
