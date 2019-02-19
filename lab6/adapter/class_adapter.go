package adapter

import (
	"io"

	"github.com/AlexanderFadeev/ood/lab6/graphics"
	"github.com/AlexanderFadeev/ood/lab6/modern_graphics"
)

type CanvasClassAdapter interface {
	graphics.Canvas
	modern_graphics.Renderer
}

type canvasClassAdapter struct {
	modern_graphics.Renderer
	prevPoint modern_graphics.Point
	color     uint32
}

func NewCanvasClassAdapter(writer io.Writer) CanvasClassAdapter {
	return &canvasClassAdapter{
		Renderer: modern_graphics.NewRenderer(writer),
	}
}

func (a *canvasClassAdapter) SetColor(color uint32) {
	a.color = color
}

func (a *canvasClassAdapter) MoveTo(x, y int) {
	a.prevPoint = modern_graphics.Point{x, y}
}

func (a *canvasClassAdapter) LineTo(x, y int) {
	point := modern_graphics.Point{x, y}
	a.DrawLine(a.prevPoint, point, rgbToRGBA(a.color))
	a.prevPoint = point
}
