package shape

import (
	"testing"

	"ood/lab4/canvas"
	"ood/lab4/point"

	"github.com/stretchr/testify/assert"
)

func TestNewRectangle(t *testing.T) {
	table := []struct{ a, b, lt, rb point.Point }{
		{point.Point{1, 2}, point.Point{3, 4}, point.Point{1, 2}, point.Point{3, 4}},
		{point.Point{3, 2}, point.Point{1, 4}, point.Point{1, 2}, point.Point{3, 4}},
		{point.Point{1, 4}, point.Point{3, 2}, point.Point{1, 2}, point.Point{3, 4}},
		{point.Point{3, 4}, point.Point{1, 2}, point.Point{1, 2}, point.Point{3, 4}},
	}

	for _, row := range table {
		rect := NewRectangle(row.a, row.b)
		assert.Equal(t, row.lt, rect.GetLeftTop())
		assert.Equal(t, row.rb, rect.GetRightBottom())
	}

}

func TestDrawRectangle(t *testing.T) {
	pa := point.Point{-1, -5}
	pb := point.Point{10, -5}
	pc := point.Point{10, 7}
	pd := point.Point{-1, 7}

	canvasMock := new(canvas.MockCanvas)
	canvasMock.On("DrawLine", pa, pb).Return()
	canvasMock.On("DrawLine", pb, pc).Return()
	canvasMock.On("DrawLine", pc, pd).Return()
	canvasMock.On("DrawLine", pd, pa).Return()

	rectangle := NewRectangle(pb, pd)
	rectangle.Draw(canvasMock)

	canvasMock.AssertExpectations(t)
	canvasMock.AssertNumberOfCalls(t, "DrawLine", 4)
}
