package canvas

import "ood/lab4/point"

type line struct {
	from point.Point
	to   point.Point
}

func (l line) draw(c pointCanvas) {
	c.drawPoint(l.from)
	c.drawPoint(l.to)
	l.drawImpl(c, l.from, l.to)
}

func (l line) drawImpl(c pointCanvas, from, to point.Point) {
	mid := point.Point{
		X: (from.X + to.X) / 2,
		Y: (from.Y + to.Y) / 2,
	}
	c.drawPoint(mid)

	if from.DistTo(to) < pixelRadius {
		return
	}

	l.drawImpl(c, from, mid)
	l.drawImpl(c, mid, to)

}
