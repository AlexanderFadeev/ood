package shape

import (
	"fmt"

	"github.com/AlexanderFadeev/ood/lab4/canvas"
)

type Shape interface {
	fmt.Stringer
	canvas.Drawable
}
