package shape

import (
	"image/color"
)

type groupLineStyle struct {
	shapes *shapes
}

func (gls *groupLineStyle) IsEnabled() *bool {
	if gls.shapes.GetShapesCount() == 0 {
		return nil
	}

	result := gls.shapes.GetShape(0).GetLineStyle().IsEnabled()
	for index := range gls.shapes.Vector {
		val := gls.shapes.GetShape(index).GetLineStyle().IsEnabled()
		if val == nil || *result != *val {
			return nil
		}
	}

	return result
}

func (gls *groupLineStyle) Enable(enabled bool) {
	for _, s := range gls.shapes.Vector {
		s.(Shape).GetLineStyle().Enable(enabled)
	}
}

func (gls *groupLineStyle) GetColor() color.Color {
	if gls.shapes.GetShapesCount() == 0 {
		return nil
	}

	result := gls.shapes.GetShape(0).GetLineStyle().GetColor()
	for _, s := range gls.shapes.Vector {
		val := s.(Shape).GetLineStyle().GetColor()
		if val == nil || !colorsAreEqual(result, val) {
			return nil
		}
	}
	return result
}

func (gls *groupLineStyle) SetColor(color color.Color) {
	for index := range gls.shapes.Vector {
		gls.shapes.GetShape(index).GetLineStyle().SetColor(color)
	}
}

func (gls *groupLineStyle) GetWidth() *float64 {
	if gls.shapes.GetShapesCount() == 0 {
		return nil
	}

	result := gls.shapes.GetShape(0).GetLineStyle().GetWidth()
	for _, s := range gls.shapes.Vector {
		val := s.(Shape).GetLineStyle().GetWidth()
		if val == nil || *result != *val {
			return nil
		}
	}
	return result
}

func (gls *groupLineStyle) SetWidth(width float64) {
	for index := range gls.shapes.Vector {
		gls.shapes.GetShape(index).GetLineStyle().SetWidth(width)
	}
}
