package style

import (
	"image/color"
)

type FillStyleEnumerator interface {
	Count() int
	FillStyle(int) FillStyle
}

type compoundFillStyle struct {
	enum FillStyleEnumerator
}

func NewCompoundFillStyle(enum FillStyleEnumerator) FillStyle {
	return &compoundFillStyle{
		enum: enum,
	}
}

func (fs *compoundFillStyle) IsEnabled() *bool {
	if fs.enum.Count() == 0 {
		return nil
	}

	result := fs.enum.FillStyle(0).IsEnabled()
	for index := 0; index < fs.enum.Count(); index++ {
		s := fs.enum.FillStyle(index)
		val := s.IsEnabled()
		if val == nil || *result != *val {
			return nil
		}
	}

	return result
}

func (fs *compoundFillStyle) Enable(enabled bool) {
	for index := 0; index < fs.enum.Count(); index++ {
		s := fs.enum.FillStyle(index)
		s.Enable(enabled)
	}
}

func (fs *compoundFillStyle) GetColor() color.Color {
	if fs.enum.Count() == 0 {
		return nil
	}

	result := fs.enum.FillStyle(0).GetColor()
	for index := 0; index < fs.enum.Count(); index++ {
		s := fs.enum.FillStyle(index)
		val := s.GetColor()
		if val == nil || !colorsAreEqual(result, val) {
			return nil
		}
	}

	return result
}

func (fs *compoundFillStyle) SetColor(color color.Color) {
	for index := 0; index < fs.enum.Count(); index++ {
		s := fs.enum.FillStyle(index)
		s.SetColor(color)
	}
}
