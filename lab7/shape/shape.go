package shape

import (
	"ood/lab7/rect"
	"ood/lab7/style"
)

type Shape interface {
	GetFrame() *rect.Rect
	SetFrame(rect.Rect)

	GetLineStyle() style.LineStyle
	GetFillStyle() style.FillStyle

	GetGroup() Group

	Accept(Visitor)
}
