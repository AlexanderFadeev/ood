package adapter

import (
	"ood/lab6/graphics"
	"ood/lab6/modern_graphics"
)

type objectAdapter struct {
	adaptee   modern_graphics.Renderer
	prevPoint modern_graphics.Point
	color     uint32
}

func NewObjectAdapter(renderer modern_graphics.Renderer) graphics.Canvas {
	return &objectAdapter{
		adaptee: renderer,
	}
}

func (a *objectAdapter) SetColor(color uint32) {
	a.color = color
}

func (a *objectAdapter) MoveTo(x, y int) {
	a.prevPoint = modern_graphics.Point{x, y}
}

func (a *objectAdapter) LineTo(x, y int) {
	point := modern_graphics.Point{x, y}
	a.adaptee.DrawLine(a.prevPoint, point, rgbToRGBA(a.color))
	a.prevPoint = point
}
