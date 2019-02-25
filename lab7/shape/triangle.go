package shape

import (
	"github.com/AlexanderFadeev/ood/lab7/canvas"
	"github.com/AlexanderFadeev/ood/lab7/point"
	"github.com/AlexanderFadeev/ood/lab7/rect"
)

type Triangle struct {
	leaf
	styles
	frame

	verticesRelPos [3]point.Point
}

func NewTriangle(vertices [3]point.Point) Shape {
	frame := rect.New(vertices[0], vertices[1]).Extend(vertices[2])
	relPos := vertices
	for index := range relPos {
		relPos[index].X -= frame.LeftTop.X
		relPos[index].Y -= frame.LeftTop.Y
		relPos[index].X /= frame.Width()
		relPos[index].Y /= frame.Height()
	}

	return &Triangle{
		frame:          *NewFrame(frame),
		styles:         *NewStyles(),
		verticesRelPos: relPos,
	}
}

func (t *Triangle) Draw(c canvas.Canvas) {
	t.styles.apply(c)
	c.MoveTo(t.GetVertex(0))
	c.LineTo(t.GetVertex(1))
	c.LineTo(t.GetVertex(2))
	c.LineTo(t.GetVertex(0))
}

func (t *Triangle) GetVertex(index uint) point.Point {
	return t.frame.LeftTop.Shift(
		t.Dimensions().PairwiseProduct(point.Vector(t.verticesRelPos[index])),
	)
}
