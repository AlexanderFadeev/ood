package style

import (
	"image/color"
)

type style interface {
	IsEnabled() *bool
	Enable(bool)

	GetColor() color.Color
	SetColor(color.Color)
}

type FillStyle interface {
	style
}

type LineStyle interface {
	style

	GetWidth() *float64
	SetWidth(float64)
}

type styleImpl struct {
	enabled bool
	color   color.Color
}

func newStyleImpl() styleImpl {
	return styleImpl{
		enabled: true,
		color:   color.Transparent,
	}
}

func (s *styleImpl) IsEnabled() *bool {
	return &s.enabled
}

func (s *styleImpl) Enable(enabled bool) {
	s.enabled = enabled
}

func (s *styleImpl) GetColor() color.Color {
	return s.color
}

func (s *styleImpl) SetColor(color color.Color) {
	s.color = color
}
