package style

import (
	"image/color"
)

type LineStyleEnumerator interface {
	Count() int
	LineStyle(int) LineStyle
	LineStyles() []LineStyle
}

type compoundLineStyle struct {
	enum LineStyleEnumerator
}

func NewCompoundLineStyle(enum LineStyleEnumerator) LineStyle {
	return &compoundLineStyle{
		enum: enum,
	}
}

func (ls *compoundLineStyle) IsEnabled() *bool {
	if ls.enum.Count() == 0 {
		return nil
	}

	result := ls.enum.LineStyle(0).IsEnabled()
	for _, s := range ls.enum.LineStyles() {
		val := s.IsEnabled()
		if val == nil || *result != *val {
			return nil
		}
	}

	return result
}

func (ls *compoundLineStyle) Enable(enabled bool) {
	for _, s := range ls.enum.LineStyles() {
		s.Enable(enabled)
	}
}

func (ls *compoundLineStyle) GetColor() color.Color {
	if ls.enum.Count() == 0 {
		return nil
	}

	result := ls.enum.LineStyle(0).GetColor()
	for _, s := range ls.enum.LineStyles() {
		val := s.GetColor()
		if val == nil || !colorsAreEqual(result, val) {
			return nil
		}
	}
	return result
}

func (ls *compoundLineStyle) SetColor(color color.Color) {
	for _, s := range ls.enum.LineStyles() {
		s.SetColor(color)
	}
}

func (ls *compoundLineStyle) GetWidth() *float64 {
	if ls.enum.Count() == 0 {
		return nil
	}

	result := ls.enum.LineStyle(0).GetWidth()
	for _, s := range ls.enum.LineStyles() {
		val := s.GetWidth()
		if val == nil || *result != *val {
			return nil
		}
	}
	return result
}

func (ls *compoundLineStyle) SetWidth(width float64) {
	for _, s := range ls.enum.LineStyles() {
		s.SetWidth(width)
	}
}
