package adapter

import (
	"bytes"
	"testing"

	"ood/lab6/modern_graphics"
	"ood/lab6/shape_drawing"

	"github.com/stretchr/testify/assert"
)

func TestObjectAdapter(t *testing.T) {
	buf := new(bytes.Buffer)
	renderer := modern_graphics.NewRenderer(buf)
	adapter := NewObjectAdapter(renderer)

	renderer.BeginDraw()
	painter := shape_drawing.NewPainter(adapter)
	painter.Draw(rect)
	renderer.EndDraw()

	assert.Equal(t, expected, buf.String())
}
