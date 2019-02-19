package shape

import (
	"testing"

	"github.com/AlexanderFadeev/ood/lab4/canvas"
	"github.com/AlexanderFadeev/ood/lab4/color"
	"github.com/AlexanderFadeev/ood/lab4/point"

	"github.com/stretchr/testify/assert"
)

func TestNewEllipse(t *testing.T) {
	table := []struct {
		valid  bool
		c      point.Point
		hr, vr float64
		color  color.Color
	}{
		{true, point.Point{1, 1}, 1, 1, color.Red},
		{false, point.Point{1, 1}, -1, 1, color.Green},
		{false, point.Point{1, 1}, 1, -1, color.Blue},
		{false, point.Point{1, 1}, -1, -1, color.Pink},
	}

	for _, row := range table {
		ellipse, err := NewEllipse(row.c, row.hr, row.vr, row.color)
		if !row.valid {
			assert.NotNil(t, err)
			continue
		}

		assert.Nil(t, err)
		assert.Equal(t, row.c, ellipse.GetCenter())
		assert.Equal(t, row.hr, ellipse.GetHorizontalRadius())
		assert.Equal(t, row.vr, ellipse.GetVerticalRadius())
		assert.Equal(t, row.color, ellipse.GetColor())
	}
}

func TestDrawEllipse(t *testing.T) {
	center := point.Point{-1, -3}
	const width = 10.0
	const height = 6.0
	topLeft := point.Point{center.X - width/2, center.Y - height/2}
	const col = color.Yellow

	canvasMock := new(canvas.MockCanvas)
	canvasMock.On("DrawEllipse", topLeft, width, height).Return()
	canvasMock.On("SetColor", col).Return()

	ellipse, _ := NewEllipse(center, width/2, height/2, col)
	ellipse.Draw(canvasMock)

	canvasMock.AssertExpectations(t)
	canvasMock.AssertNumberOfCalls(t, "DrawEllipse", 1)
	canvasMock.AssertNumberOfCalls(t, "SetColor", 1)
}
