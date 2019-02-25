package shape

import (
	"github.com/AlexanderFadeev/ood/lab7/canvas"
	"github.com/AlexanderFadeev/ood/lab7/style"
)

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

func (s *styles) apply(c canvas.Canvas) {
	if *s.fillStyle.IsEnabled() {
		c.SetFillColor(s.fillStyle.GetColor())
	} else {
		c.SetFillColor(nil)
	}

	if *s.lineStyle.IsEnabled() {
		c.SetLineColor(s.lineStyle.GetColor())
		c.SetLineWidth(*s.lineStyle.GetWidth())
	} else {
		c.SetLineColor(nil)
	}
}
