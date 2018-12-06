package shape

import (
	"ood/lab7/rect"
	"ood/lab7/style"
)

type Group interface {
	Shape

	GetShapesCount() int
	GetShape(index int) Shape
	InsertShape(shape Shape, index int) bool
	RemoveShape(index int)

	getShapes() *shapes
}

type group struct {
	*shapes
}

func NewGroup() Group {
	return &group{
		shapes: new(shapes),
	}
}

func (g *group) GetGroup() Group {
	return g
}

func (g *group) GetFrame() *rect.Rect {
	if g.GetShapesCount() == 0 {
		return nil
	}

	var result *rect.Rect
	for _, s := range g.shapes.Vector {
		other := s.(Shape).GetFrame()
		if other == nil {
			continue
		}

		if result == nil {
			result = other
			continue
		}

		val := result.Outersect(*other)
		result = &val
	}

	return result
}

func (g *group) SetFrame(newFrame rect.Rect) {
	oldFrame := g.GetFrame()
	if oldFrame == nil {
		return
	}

	for index := range g.shapes.Vector {
		frame := g.shapes.GetShape(index).GetFrame()
		if frame == nil {
			continue
		}

		*frame = frame.Resize(*oldFrame, newFrame)
	}
}

func (g *group) GetLineStyle() style.LineStyle {
	return &groupLineStyle{
		shapes: g.shapes,
	}
}

func (g *group) GetFillStyle() style.FillStyle {
	return &fillStyle{
		shapes: g.shapes,
	}
}

func (g *group) Accept(v Visitor) {
	for _, s := range g.shapes.Vector {
		s.(Shape).Accept(v)
	}
}

func (g *group) getShapes() *shapes {
	return g.shapes
}
