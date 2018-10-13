package shape_drawing

import "ood/lab6/graphics"

type Drawable interface {
	Draw(graphics.Canvas)
}
