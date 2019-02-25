package style

import (
	"image/color"
)

type LineStyleEnumerator interface {
	Count() int
	LineStyle(int) LineStyle
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
	for index := 0; index < ls.enum.Count(); index++ {
		s := ls.enum.LineStyle(index)
		val := s.IsEnabled()
		if val == nil || *result != *val {
			return nil
		}
	}

	return result
}

func (ls *compoundLineStyle) Enable(enabled bool) {
	for index := 0; index < ls.enum.Count(); index++ {
		s := ls.enum.LineStyle(index)
		s.Enable(enabled)
	}
}

func (ls *compoundLineStyle) GetColor() color.Color {
	if ls.enum.Count() == 0 {
		return nil
	}

	result := ls.enum.LineStyle(0).GetColor()
	for index := 0; index < ls.enum.Count(); index++ {
		s := ls.enum.LineStyle(index)
		val := s.GetColor()
		if val == nil || !colorsAreEqual(result, val) {
			return nil
		}
	}
	return result
}

func (ls *compoundLineStyle) SetColor(color color.Color) {
	for index := 0; index < ls.enum.Count(); index++ {
		s := ls.enum.LineStyle(index)
		s.SetColor(color)
	}
}

func (ls *compoundLineStyle) GetWidth() *float64 {
	if ls.enum.Count() == 0 {
		return nil
	}

	result := ls.enum.LineStyle(0).GetWidth()
	for index := 0; index < ls.enum.Count(); index++ {
		s := ls.enum.LineStyle(index)
		val := s.GetWidth()
		if val == nil || *result != *val {
			return nil
		}
	}
	return result
}

func (ls *compoundLineStyle) SetWidth(width float64) {
	for index := 0; index < ls.enum.Count(); index++ {
		s := ls.enum.LineStyle(index)
		s.SetWidth(width)
	}
}
