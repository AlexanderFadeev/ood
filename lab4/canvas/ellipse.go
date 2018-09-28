package canvas

import (
	"math"

	"ood/lab4/point"
)

type ellipse struct {
	center           point.Point
	horizontalRadius float64
	verticalRadius   float64
}

func (e ellipse) draw(c pointCanvas) {
	c.drawPoint(e.pointFromAngle(0))
	c.drawPoint(e.pointFromAngle(math.Pi))
	e.drawImpl(c, 0, math.Pi)
	e.drawImpl(c, math.Pi, 2*math.Pi)
}

func (e ellipse) drawImpl(c pointCanvas, fromAngle, toAngle float64) {
	midAngle := (fromAngle + toAngle) / 2
	mid := e.pointFromAngle(midAngle)
	c.drawPoint(mid)

	from := e.pointFromAngle(fromAngle)
	to := e.pointFromAngle(toAngle)
	if from.DistTo(to) < pixelRadius {
		return
	}

	e.drawImpl(c, fromAngle, midAngle)
	e.drawImpl(c, midAngle, toAngle)
}

func (e ellipse) pointFromAngle(angle float64) point.Point {
	x := e.center.X + math.Cos(angle)*e.horizontalRadius
	y := e.center.Y - math.Sin(angle)*e.verticalRadius
	return point.Point{
		X: x,
		Y: y,
	}
}
