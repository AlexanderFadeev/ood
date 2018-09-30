package shape

import (
	"math"
	"testing"

	"ood/lab4/canvas"
	"ood/lab4/point"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewRegularPolygon(t *testing.T) {
	table := []struct {
		valid bool
		v     uint
		c     point.Point
		r     float64
	}{
		{true, 42, point.Point{1, 1}, 1},
		{false, 2, point.Point{1, 1}, 1},
		{false, 42, point.Point{1, 1}, -1},
	}

	for _, row := range table {
		polygon, err := NewRegularPolygon(row.v, row.c, row.r)
		if !row.valid {
			assert.NotNil(t, err)
			continue
		}

		assert.Nil(t, err)
		assert.Equal(t, row.v, polygon.GetVerticesCount())
		assert.Equal(t, row.c, polygon.GetCenter())
		assert.Equal(t, row.r, polygon.GetRadius())
	}
}

func float64AlmostEqual(a, b float64) bool {
	const eps = 1e-6
	return math.Abs(a-b) < eps
}

func pointsAlmostEqual(a point.Point) func(point.Point) bool {
	return func(b point.Point) bool {
		return float64AlmostEqual(a.X, b.X) && float64AlmostEqual(a.Y, b.Y)
	}
}

func TestDrawRegularPolygon(t *testing.T) {
	const radius = 42.0
	const vertices = 8

	polygon, _ := NewRegularPolygon(vertices, point.Point{radius, radius}, radius)

	var p [vertices]point.Point
	for i := 0; i < vertices; i++ {
		angle := float64(i) / vertices * 2 * math.Pi
		p[i] = point.Point{
			X: radius * (1 + math.Cos(angle)),
			Y: radius * (1 - math.Sin(angle)),
		}
	}

	canvasMock := new(canvas.MockCanvas)
	for i := 0; i < vertices; i++ {
		canvasMock.On("DrawLine",
			mock.MatchedBy(pointsAlmostEqual(p[i])),
			mock.MatchedBy(pointsAlmostEqual(p[(i+1)%vertices])),
		).Return()
	}

	polygon.Draw(canvasMock)
	canvasMock.AssertExpectations(t)
	canvasMock.AssertNumberOfCalls(t, "DrawLine", 8)
}
