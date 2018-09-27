package shape

import "github.com/pkg/errors"

type Ellipse struct {
	center           Point
	horizontalRadius float64
	verticalRadius   float64
}

func NewEllipse(center Point, horizontalRadius, verticalRadius float64) (*Ellipse, error) {
	if horizontalRadius < 0 {
		return nil, errors.New("Negative horizontal radius value")
	}
	if verticalRadius < 0 {
		return nil, errors.New("Negative vertical radius value")
	}

	return &Ellipse{
		center:           center,
		horizontalRadius: horizontalRadius,
		verticalRadius:   verticalRadius,
	}, nil
}

func (e Ellipse) GetCenter() Point {
	return e.center
}

func (e Ellipse) GetHorizontalRadius() float64 {
	return e.horizontalRadius
}

func (e Ellipse) GetVerticalRadius() float64 {
	return e.verticalRadius
}
