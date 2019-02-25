package adapter

import (
	"bytes"
	"testing"

	"github.com/AlexanderFadeev/ood/lab6/modern_graphics"
	"github.com/AlexanderFadeev/ood/lab6/shape_drawing"

	"github.com/stretchr/testify/assert"
)

func TestCanvasObjectAdapter(t *testing.T) {
	buf := new(bytes.Buffer)
	renderer := modern_graphics.NewRenderer(buf)
	adapter := NewCanvasObjectAdapter(renderer)

	renderer.BeginDraw()
	painter := shape_drawing.NewPainter(adapter)
	painter.Draw(rect)
	renderer.EndDraw()

	assert.Equal(t, expected, buf.String())
}
