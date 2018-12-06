package shape

type Visitor interface {
	VisitEllipse(*Ellipse)
	VisitRectangle(*Rectangle)
	VisitTriangle(*Triangle)
}
