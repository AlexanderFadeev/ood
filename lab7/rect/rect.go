package rect

import (
	"math"

	"ood/lab7/point"
)

type Rect struct {
	LeftTop     point.Point
	RightBottom point.Point
}

func New(a, b point.Point) Rect {
	return Rect{
		LeftTop: point.Point{
			X: math.Min(a.X, b.X),
			Y: math.Min(a.Y, b.Y),
		},
		RightBottom: point.Point{
			X: math.Max(a.X, b.X),
			Y: math.Max(a.Y, b.Y),
		},
	}
}

func (r Rect) Extend(p point.Point) Rect {
	return Rect{
		LeftTop: point.Point{
			X: math.Min(r.LeftTop.X, p.X),
			Y: math.Min(r.LeftTop.Y, p.Y),
		},
		RightBottom: point.Point{
			X: math.Max(r.RightBottom.X, p.X),
			Y: math.Max(r.RightBottom.Y, p.Y),
		},
	}
}

func (r Rect) Outersect(other Rect) Rect {
	return r.Extend(other.LeftTop).Extend(other.RightBottom)
}

func (r Rect) Resize(from, to Rect) Rect {
	delta1 := point.Vector(from.LeftTop).Negative()
	delta2 := point.Vector(to.LeftTop)
	scale := to.Dimensions().PairwiseDivision(from.Dimensions())
	return r.shift(delta1).scale(scale).shift(delta2)
}

func (r *Rect) Width() float64 {
	return r.RightBottom.X - r.LeftTop.X
}

func (r *Rect) Height() float64 {
	return r.RightBottom.Y - r.LeftTop.Y
}

func (r *Rect) Dimensions() point.Vector {
	return point.NewVector(r.LeftTop, r.RightBottom)
}

func (r Rect) shift(vec point.Vector) Rect {
	return Rect{
		LeftTop:     r.LeftTop.Shift(vec),
		RightBottom: r.RightBottom.Shift(vec),
	}
}

//Scales relative to (0, 0)
func (r Rect) scale(vec point.Vector) Rect {
	return Rect{
		LeftTop:     point.Point(point.Vector(r.LeftTop).PairwiseProduct(vec)),
		RightBottom: point.Point(point.Vector(r.RightBottom).PairwiseProduct(vec)),
	}
}
