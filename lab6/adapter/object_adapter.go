package adapter

import (
	"ood/lab6/graphics"
	"ood/lab6/modern_graphics"
)

type canvasObjectAdapter struct {
	adaptee   modern_graphics.Renderer
	prevPoint modern_graphics.Point
	color     uint32
}

func NewCanvasObjectAdapter(renderer modern_graphics.Renderer) graphics.Canvas {
	return &canvasObjectAdapter{
		adaptee: renderer,
	}
}

func (a *canvasObjectAdapter) SetColor(color uint32) {
	a.color = color
}

func (a *canvasObjectAdapter) MoveTo(x, y int) {
	a.prevPoint = modern_graphics.Point{x, y}
}

func (a *canvasObjectAdapter) LineTo(x, y int) {
	point := modern_graphics.Point{x, y}
	a.adaptee.DrawLine(a.prevPoint, point, rgbToRGBA(a.color))
	a.prevPoint = point
}
