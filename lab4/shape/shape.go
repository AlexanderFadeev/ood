package shape

import "ood/lab4/canvas"

type Shape interface {
	Draw(canvas.Canvas)
}
