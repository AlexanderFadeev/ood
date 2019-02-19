package shape

import (
	"testing"

	"github.com/AlexanderFadeev/ood/lab4/canvas"
	"github.com/AlexanderFadeev/ood/lab4/color"
	"github.com/AlexanderFadeev/ood/lab4/point"

	"github.com/stretchr/testify/assert"
)

func TestNewTriangle(t *testing.T) {
	table := []struct {
		a, b, c point.Point
		color   color.Color
	}{
		{point.Point{1, 2}, point.Point{3, 4}, point.Point{5, 6}, color.Red},
		{point.Point{-1, -2}, point.Point{-3, -4}, point.Point{-5, -6}, color.Green},
		{point.Point{1, 2}, point.Point{-3, 4}, point.Point{-5, -6}, color.Blue},
	}

	for _, row := range table {
		triangle := NewTriangle(row.a, row.b, row.c, row.color)
		assert.Equal(t, row.a, triangle.GetVertexA())
		assert.Equal(t, row.b, triangle.GetVertexB())
		assert.Equal(t, row.c, triangle.GetVertexC())
		assert.Equal(t, row.color, triangle.GetColor())
	}
}

func TestDrawTriangle(t *testing.T) {
	pa := point.Point{1, 1}
	pb := point.Point{5, 5}
	pc := point.Point{1, 5}
	const col = color.Black

	canvasMock := new(canvas.MockCanvas)
	canvasMock.On("DrawLine", pa, pb).Return()
	canvasMock.On("DrawLine", pb, pc).Return()
	canvasMock.On("DrawLine", pc, pa).Return()
	canvasMock.On("SetColor", col).Return()

	triangle := NewTriangle(pa, pb, pc, col)
	triangle.Draw(canvasMock)

	canvasMock.AssertExpectations(t)
	canvasMock.AssertNumberOfCalls(t, "DrawLine", 3)
	canvasMock.AssertNumberOfCalls(t, "SetColor", 1)
}
