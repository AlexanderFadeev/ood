package shape

import (
	"image/color"
)

type fillStyle struct {
	shapes *shapes
}

func (fs *fillStyle) IsEnabled() *bool {
	if fs.shapes.GetShapesCount() == 0 {
		return nil
	}

	result := fs.shapes.GetShape(0).GetFillStyle().IsEnabled()
	for index := range fs.shapes.Vector {
		val := fs.shapes.GetShape(index).GetFillStyle().IsEnabled()
		if val == nil || *result != *val {
			return nil
		}
	}

	return result
}

func (fs *fillStyle) Enable(enabled bool) {
	for _, s := range fs.shapes.Vector {
		s.(Shape).GetFillStyle().Enable(enabled)
	}
}

func (fs *fillStyle) GetColor() color.Color {
	if fs.shapes.GetShapesCount() == 0 {
		return nil
	}

	result := fs.shapes.GetShape(0).GetFillStyle().GetColor()
	for _, s := range fs.shapes.Vector {
		val := s.(Shape).GetFillStyle().GetColor()
		if val == nil || !colorsAreEqual(result, val) {
			return nil
		}
	}

	return result
}

func (fs *fillStyle) SetColor(color color.Color) {
	for index := range fs.shapes.Vector {
		fs.shapes.GetShape(index).GetFillStyle().SetColor(color)
	}
}
