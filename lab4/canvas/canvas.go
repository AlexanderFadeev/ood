package canvas

import (
	"ood/lab4/color"
	"ood/lab4/point"
)

type Canvas interface {
	SetColor(color.Color)
	DrawLine(from, to point.Point)
	DrawEllipse(leftTop point.Point, width, height float64)
}
