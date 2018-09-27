package shape

import (
	"testing"

	"ood/lab4/point"

	"github.com/stretchr/testify/assert"
)

func TestNewTriangle(t *testing.T) {
	table := []struct{ a, b, c point.Point }{
		{point.Point{1, 2}, point.Point{3, 4}, point.Point{5, 6}},
		{point.Point{-1, -2}, point.Point{-3, -4}, point.Point{-5, -6}},
		{point.Point{1, 2}, point.Point{-3, 4}, point.Point{-5, -6}},
	}

	for _, row := range table {
		triangle := NewTriangle(row.a, row.b, row.c)
		assert.Equal(t, row.a, triangle.GetVertexA())
		assert.Equal(t, row.b, triangle.GetVertexB())
		assert.Equal(t, row.c, triangle.GetVertexC())
	}
}
