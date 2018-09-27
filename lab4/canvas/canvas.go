package canvas

import (
	"ood/lab4/color"
	"ood/lab4/shape"
)

type Canvas interface {
	SetColor(color.Color)
	DrawLine(from, to shape.Point)
	DrawEllipse(leftTop shape.Point, width, height float64)
}
