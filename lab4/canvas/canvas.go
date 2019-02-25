package canvas

import (
	"image"
	"io"

	"github.com/AlexanderFadeev/ood/browser_display"
	"github.com/AlexanderFadeev/ood/lab4/color"
	"github.com/AlexanderFadeev/ood/lab4/point"

	"golang.org/x/image/bmp"
)

const pixelRadius = 0.5

type Drawable interface {
	Draw(canvas)
}

type Canvas interface {
	SetColor(color.Color)
	DrawLine(from, to point.Point)
	DrawEllipse(leftTop point.Point, width, height float64)
}

type DisplayableCanvas interface {
	Canvas
	Display()
}

type pointCanvas interface {
	drawPoint(point.Point)
}

type canvas struct {
	color color.Color
	img   *image.RGBA
}

func New(width, height uint) DisplayableCanvas {
	rect := image.Rect(0, 0, int(width), int(height))
	return &canvas{
		img: image.NewRGBA(rect),
	}
}

func (c *canvas) SetColor(color color.Color) {
	c.color = color
}

func (c *canvas) DrawLine(from, to point.Point) {
	line := line{
		from: from,
		to:   to,
	}
	line.draw(c)
}

func (c *canvas) DrawEllipse(leftTop point.Point, width, height float64) {
	center := point.Point{
		X: leftTop.X + width/2,
		Y: leftTop.Y + height/2,
	}
	ellipse := ellipse{
		center:           center,
		horizontalRadius: width / 2,
		verticalRadius:   height / 2,
	}
	ellipse.draw(c)
}

func (pc *canvas) drawPoint(p point.Point) {
	pc.img.Set(int(p.X), int(p.Y), pc.color)
}

func (c *canvas) Display() {
	browser_display.DisplayInBrowser(func(w io.Writer) {
		bmp.Encode(w, c.img)
	}, "image/bmp")
}
