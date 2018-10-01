package shape

import "ood/lab4/color"

type shapeColor color.Color

func (c shapeColor) GetColor() color.Color {
	return color.Color(c)
}
