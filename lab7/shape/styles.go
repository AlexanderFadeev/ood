package shape

import "ood/lab7/style"

type styles struct {
	lineStyle style.LineStyle
	fillStyle style.FillStyle
}

func NewStyles() *styles {
	return &styles{
		lineStyle: style.NewLineStyle(),
		fillStyle: style.NewFillStyle(),
	}
}

func (s *styles) GetLineStyle() style.LineStyle {
	return s.lineStyle
}

func (s *styles) GetFillStyle() style.FillStyle {
	return s.fillStyle
}
