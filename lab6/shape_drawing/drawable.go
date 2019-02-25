package shape_drawing

import "github.com/AlexanderFadeev/ood/lab6/graphics"

type Drawable interface {
	Draw(graphics.Canvas)
}
