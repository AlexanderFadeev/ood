package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRectangle(t *testing.T) {
	table := []struct{ a, b, lt, rb Point }{
		{Point{1, 2}, Point{3, 4}, Point{1, 2}, Point{3, 4}},
		{Point{3, 2}, Point{1, 4}, Point{1, 2}, Point{3, 4}},
		{Point{1, 4}, Point{3, 2}, Point{1, 2}, Point{3, 4}},
		{Point{3, 4}, Point{1, 2}, Point{1, 2}, Point{3, 4}},
	}

	for _, row := range table {
		rect := NewRectangle(row.a, row.b)
		assert.Equal(t, row.lt, rect.GetLeftTop())
		assert.Equal(t, row.rb, rect.GetRightBottom())
	}

}

//func TestDrawRectangle(t *testing.T) {
//	canvas := canvas.NewMock()
//	rect := NewRectangle({1, 2}, {3, 4})
//}
