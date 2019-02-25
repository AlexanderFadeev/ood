package shape

import (
	"fmt"

	"github.com/AlexanderFadeev/ood/lab4/canvas"
	"github.com/AlexanderFadeev/ood/lab4/color"
	"github.com/AlexanderFadeev/ood/lab4/point"

	"github.com/pkg/errors"
)

type Ellipse struct {
	shapeColor

	center           point.Point
	horizontalRadius float64
	verticalRadius   float64
}

func NewEllipse(center point.Point, horizontalRadius, verticalRadius float64, color color.Color) (*Ellipse, error) {
	if horizontalRadius < 0 {
		return nil, errors.New("Negative horizontal radius value")
	}
	if verticalRadius < 0 {
		return nil, errors.New("Negative vertical radius value")
	}

	return &Ellipse{
		shapeColor:       shapeColor(color),
		center:           center,
		horizontalRadius: horizontalRadius,
		verticalRadius:   verticalRadius,
	}, nil
}

func (e Ellipse) GetCenter() point.Point {
	return e.center
}

func (e Ellipse) GetHorizontalRadius() float64 {
	return e.horizontalRadius
}

func (e Ellipse) GetVerticalRadius() float64 {
	return e.verticalRadius
}

func (e Ellipse) Draw(canvas canvas.Canvas) {
	canvas.SetColor(e.GetColor())
	topLeft := point.Point{e.center.X - e.horizontalRadius, e.center.Y - e.verticalRadius}
	canvas.DrawEllipse(topLeft, e.horizontalRadius*2, e.verticalRadius*2)
}

func (e Ellipse) String() string {
	return fmt.Sprintf("%s ellipse: O=%s, HR=%.1f, VR=%.1f", e.GetColor(), e.center, e.horizontalRadius, e.verticalRadius)
}
