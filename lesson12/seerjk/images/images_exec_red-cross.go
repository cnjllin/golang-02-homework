package images

import (
	"golang.org/x/tour/pic"
	"image/color"
	"image"
	"math"
)

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
	// 绘制红色十字
	halfX := im.Width / 2
	halfY := im.Height / 2
	crossWidth := int(math.Min(float64(im.Width), float64(im.Height)) / 5)
	if (x >= halfX-crossWidth && x <= halfX+crossWidth) ||
		y >= halfY-crossWidth && y <= halfY+crossWidth {
		return color.RGBA{255, 0, 0, 255}
	} else {
		return color.RGBA{255, 255, 255, 255}
	}
}

func main() {
	m := Image{200, 200}
	pic.ShowImage(&m)
}
