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

func (g *group) Count() int {
	return g.GetShapesCount()
}

func (g *group) GetLineStyle() style.LineStyle {
	return style.NewCompoundLineStyle(g)
}

func (g *group) LineStyle(index int) style.LineStyle {
	return g.shapes.GetShape(index).GetLineStyle()
}

func (g *group) LineStyles() []style.LineStyle {
	var result []style.LineStyle
	for _, shape := range g.Vector {
		result = append(result, shape.(Shape).GetLineStyle())
	}
	return result
}

func (g *group) GetFillStyle() style.FillStyle {
	return style.NewCompoundFillStyle(g)
}

func (g *group) FillStyle(index int) style.FillStyle {
	return g.shapes.GetShape(index).GetFillStyle()
}

func (g *group) FillStyles() []style.FillStyle {
	var result []style.FillStyle
	for _, shape := range g.Vector {
		result = append(result, shape.(Shape).GetFillStyle())
	}
	return result
}

func (g *group) Accept(v Visitor) {
	for _, s := range g.shapes.Vector {
		s.(Shape).Accept(v)
	}
}

func (g *group) getShapes() *shapes {
	return g.shapes
}
