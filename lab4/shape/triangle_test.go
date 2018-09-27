package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTriangle(t *testing.T) {
	table := []struct{ a, b, c Point }{
		{Point{1, 2}, Point{3, 4}, Point{5, 6}},
		{Point{-1, -2}, Point{-3, -4}, Point{-5, -6}},
		{Point{1, 2}, Point{-3, 4}, Point{-5, -6}},
	}

	for _, row := range table {
		triangle := NewTriangle(row.a, row.b, row.c)
		assert.Equal(t, row.a, triangle.GetVertexA())
		assert.Equal(t, row.b, triangle.GetVertexB())
		assert.Equal(t, row.c, triangle.GetVertexC())
	}
}
