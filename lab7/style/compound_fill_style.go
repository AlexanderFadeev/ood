package style

import (
	"image/color"
)

type FillStyleEnumerator interface {
	Count() int
	FillStyle(int) FillStyle
	FillStyles() []FillStyle
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
	for _, s := range fs.enum.FillStyles() {
		val := s.IsEnabled()
		if val == nil || *result != *val {
			return nil
		}
	}

	return result
}

func (fs *compoundFillStyle) Enable(enabled bool) {
	for _, s := range fs.enum.FillStyles() {
		s.Enable(enabled)
	}
}

func (fs *compoundFillStyle) GetColor() color.Color {
	if fs.enum.Count() == 0 {
		return nil
	}

	result := fs.enum.FillStyle(0).GetColor()
	for _, s := range fs.enum.FillStyles() {
		val := s.GetColor()
		if val == nil || !colorsAreEqual(result, val) {
			return nil
		}
	}

	return result
}

func (fs *compoundFillStyle) SetColor(color color.Color) {
	for _, s := range fs.enum.FillStyles() {
		s.SetColor(color)
	}
}
