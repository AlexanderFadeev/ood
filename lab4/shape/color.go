package shape

import "github.com/AlexanderFadeev/ood/lab4/color"

type shapeColor color.Color

func (c shapeColor) GetColor() color.Color {
	return color.Color(c)
}
