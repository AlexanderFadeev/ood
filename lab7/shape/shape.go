package shape

import (
	"github.com/AlexanderFadeev/ood/lab7/canvas"
	"github.com/AlexanderFadeev/ood/lab7/rect"
	"github.com/AlexanderFadeev/ood/lab7/style"
)

type Shape interface {
	GetFrame() *rect.Rect
	SetFrame(rect.Rect)

	GetLineStyle() style.LineStyle
	GetFillStyle() style.FillStyle

	GetGroup() Group

	Draw(canvas canvas.Canvas)
}
