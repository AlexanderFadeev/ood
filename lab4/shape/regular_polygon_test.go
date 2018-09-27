package shape

import (
	"testing"

	"ood/lab4/point"

	"github.com/stretchr/testify/assert"
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
		ellipse, err := NewRegularPolygon(row.v, row.c, row.r)
		if !row.valid {
			assert.NotNil(t, err)
			continue
		}

		assert.Nil(t, err)
		assert.Equal(t, row.v, ellipse.GetVerticesCount())
		assert.Equal(t, row.c, ellipse.GetCenter())
		assert.Equal(t, row.r, ellipse.GetRadius())
	}
}
