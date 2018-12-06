package slide

import (
	"image/color"

	"ood/lab7/rect"
	"ood/lab7/shape"
)

type Slide interface {
	shape.Group

	GetBackgroundColor() color.Color
	SetBackgroundColor(color.Color)
}

type slide struct {
	shape.Group

	background           shape.Shape
	shapesWithBackground shape.Group
}

func NewSlide(rect rect.Rect) Slide {
	shapes := shape.NewGroup()
	background := shape.NewRectangle(rect)

	shapesWithBackground := shape.NewGroup()
	shapesWithBackground.InsertShape(background, 0)
	shapesWithBackground.InsertShape(shapes, 1)

	return &slide{
		Group:                shapes,
		background:           background,
		shapesWithBackground: shapesWithBackground,
	}
}

func (s *slide) SetFrame(frame rect.Rect) {
	s.shapesWithBackground.SetFrame(frame)
}

func (s *slide) GetFrame() *rect.Rect {
	return s.shapesWithBackground.GetFrame()
}

func (s *slide) Accept(v shape.Visitor) {
	s.shapesWithBackground.Accept(v)
}

func (s *slide) GetBackgroundColor() color.Color {
	return s.background.GetFillStyle().GetColor()
}

func (s *slide) SetBackgroundColor(color color.Color) {
	s.background.GetFillStyle().SetColor(color)
}
