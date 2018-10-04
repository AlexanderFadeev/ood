package shape

import (
	"fmt"

	"ood/lab4/canvas"
)

type Shape interface {
	fmt.Stringer
	canvas.Drawable
}
