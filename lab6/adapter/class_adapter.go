package adapter

import (
	"io"

	"ood/lab6/graphics"
	"ood/lab6/modern_graphics"
)

type ClassAdapter interface {
	graphics.Canvas
	modern_graphics.Renderer
}

type classAdapter struct {
	modern_graphics.Renderer
	prevPoint modern_graphics.Point
	color     uint32
}

func NewClassAdapter(writer io.Writer) ClassAdapter {
	return &classAdapter{
		Renderer: modern_graphics.NewRenderer(writer),
	}
}

func (a *classAdapter) SetColor(color uint32) {
	a.color = color
}

func (a *classAdapter) MoveTo(x, y int) {
	a.prevPoint = modern_graphics.Point{x, y}
}

func (a *classAdapter) LineTo(x, y int) {
	point := modern_graphics.Point{x, y}
	a.DrawLine(a.prevPoint, point, rgbToRGBA(a.color))
	a.prevPoint = point
}
