package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEllipse(t *testing.T) {
	table := []struct {
		valid  bool
		c      Point
		hr, vr float64
	}{
		{true, Point{1, 1}, 1, 1},
		{false, Point{1, 1}, -1, 1},
		{false, Point{1, 1}, 1, -1},
		{false, Point{1, 1}, -1, -1},
	}

	for _, row := range table {
		ellipse, err := NewEllipse(row.c, row.hr, row.vr)
		if !row.valid {
			assert.NotNil(t, err)
			continue
		}

		assert.Nil(t, err)
		assert.Equal(t, row.c, ellipse.GetCenter())
		assert.Equal(t, row.hr, ellipse.GetHorizontalRadius())
		assert.Equal(t, row.vr, ellipse.GetVerticalRadius())
	}
}
