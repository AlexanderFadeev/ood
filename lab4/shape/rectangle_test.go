package shape

import (
	"testing"

	"ood/lab4/canvas"
	"ood/lab4/color"
	"ood/lab4/point"

	"github.com/stretchr/testify/assert"
)

func TestNewRectangle(t *testing.T) {
	table := []struct {
		a, b, lt, rb point.Point
		color        color.Color
	}{
		{point.Point{1, 2}, point.Point{3, 4}, point.Point{1, 2}, point.Point{3, 4}, color.Red},
		{point.Point{3, 2}, point.Point{1, 4}, point.Point{1, 2}, point.Point{3, 4}, color.Green},
		{point.Point{1, 4}, point.Point{3, 2}, point.Point{1, 2}, point.Point{3, 4}, color.Blue},
		{point.Point{3, 4}, point.Point{1, 2}, point.Point{1, 2}, point.Point{3, 4}, color.Black},
	}

	for _, row := range table {
		rect := NewRectangle(row.a, row.b, row.color)
		assert.Equal(t, row.lt, rect.GetLeftTop())
		assert.Equal(t, row.rb, rect.GetRightBottom())
		assert.Equal(t, row.color, rect.GetColor())
	}
}

func TestDrawRectangle(t *testing.T) {
	pa := point.Point{-1, -5}
	pb := point.Point{10, -5}
	pc := point.Point{10, 7}
	pd := point.Point{-1, 7}
	const col = color.Red

	canvasMock := new(canvas.MockCanvas)
	canvasMock.On("DrawLine", pa, pb).Return()
	canvasMock.On("DrawLine", pb, pc).Return()
	canvasMock.On("DrawLine", pc, pd).Return()
	canvasMock.On("DrawLine", pd, pa).Return()
	canvasMock.On("SetColor", col).Return()

	rectangle := NewRectangle(pb, pd, col)
	rectangle.Draw(canvasMock)

	canvasMock.AssertExpectations(t)
	canvasMock.AssertNumberOfCalls(t, "DrawLine", 4)
	canvasMock.AssertNumberOfCalls(t, "SetColor", 1)
}
