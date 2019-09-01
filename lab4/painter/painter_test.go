package painter

import (
	"github.com/AlexanderFadeev/ood/lab4/canvas"
	"github.com/AlexanderFadeev/ood/lab4/picture_draft"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestPainter(t *testing.T) {
	painter := New()
	c := new(canvas.MockCanvas)

	var counter int

	d1 := new(canvas.MockDrawable)
	d2 := new(canvas.MockDrawable)
	d3 := new(canvas.MockDrawable)

	d1.On("Draw", c).Run(func(mock.Arguments) {
		assert.Equal(t, 0, counter)
		counter++
	})
	d2.On("Draw", c).Run(func(mock.Arguments) {
		assert.Equal(t, 1, counter)
		counter++
	})
	d3.On("Draw", c).Run(func(mock.Arguments) {
		assert.Equal(t, 2, counter)
		counter++
	})

	draft := picture_draft.PictureDraft{d1, d2, d3}

	painter.Paint(draft, c)
	assert.Equal(t, 3, counter)
}
