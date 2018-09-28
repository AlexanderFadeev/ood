package shape

import (
	"ood/lab4/canvas"
	"ood/lab4/point"
)

type Rectangle struct {
	leftTop     point.Point
	rightBottom point.Point
}

func NewRectangle(leftTop, rightBottom point.Point) Rectangle {
	if leftTop.X > rightBottom.X {
		leftTop.X, rightBottom.X = rightBottom.X, leftTop.X
	}
	if leftTop.Y > rightBottom.Y {
		leftTop.Y, rightBottom.Y = rightBottom.Y, leftTop.Y
	}

	return Rectangle{
		leftTop:     leftTop,
		rightBottom: rightBottom,
	}
}

func (r Rectangle) GetLeftTop() point.Point {
	return r.leftTop
}

func (r Rectangle) GetRightBottom() point.Point {
	return r.rightBottom
}

func (r Rectangle) Draw(canvas canvas.Canvas) {
	rightTop := point.Point{r.rightBottom.X, r.leftTop.Y}
	leftBottom := point.Point{r.leftTop.X, r.rightBottom.Y}

	canvas.DrawLine(r.leftTop, rightTop)
	canvas.DrawLine(rightTop, r.rightBottom)
	canvas.DrawLine(r.rightBottom, leftBottom)
	canvas.DrawLine(leftBottom, r.leftTop)
}
