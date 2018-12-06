package drawer

import (
	"ood/lab7/canvas"
	"ood/lab7/point"
	"ood/lab7/shape"
)

type Drawer interface {
	DrawShape(shape.Shape)
}

type drawer struct {
	canvas canvas.Canvas
}

func NewDrawer(canvas canvas.Canvas) Drawer {
	return &drawer{
		canvas: canvas,
	}
}

func (d *drawer) DrawShape(shape shape.Shape) {
	shape.Accept(d)
}

func (d *drawer) VisitEllipse(ellipse *shape.Ellipse) {
	d.applyStyles(ellipse)
	d.canvas.DrawEllipse(*ellipse.GetFrame())
}

func (d *drawer) VisitRectangle(rectangle *shape.Rectangle) {
	frame := rectangle.GetFrame()
	rightTop := point.Point{X: frame.RightBottom.X, Y: frame.LeftTop.Y}
	leftBottom := point.Point{X: frame.LeftTop.X, Y: frame.RightBottom.Y}

	d.applyStyles(rectangle)
	d.canvas.MoveTo(frame.LeftTop)
	d.canvas.LineTo(rightTop)
	d.canvas.LineTo(frame.RightBottom)
	d.canvas.LineTo(leftBottom)
	d.canvas.LineTo(frame.LeftTop)
}

func (d *drawer) VisitTriangle(triangle *shape.Triangle) {
	d.applyStyles(triangle)
	d.canvas.MoveTo(triangle.GetVertex(0))
	d.canvas.LineTo(triangle.GetVertex(1))
	d.canvas.LineTo(triangle.GetVertex(2))
	d.canvas.LineTo(triangle.GetVertex(0))
}

func (d *drawer) applyStyles(shape shape.Shape) {
	if *shape.GetFillStyle().IsEnabled() {
		d.canvas.SetFillColor(shape.GetFillStyle().GetColor())
	} else {
		d.canvas.SetFillColor(nil)
	}

	if *shape.GetLineStyle().IsEnabled() {
		d.canvas.SetLineColor(shape.GetLineStyle().GetColor())
		d.canvas.SetLineWidth(*shape.GetLineStyle().GetWidth())
	} else {
		d.canvas.SetLineColor(nil)
	}
}
