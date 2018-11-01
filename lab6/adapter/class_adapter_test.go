package adapter

import (
	"bytes"
	"testing"

	"ood/lab6/shape_drawing"

	"github.com/stretchr/testify/assert"
)

func TestCanvasClassAdapter(t *testing.T) {
	buf := new(bytes.Buffer)
	adapter := NewCanvasClassAdapter(buf)

	adapter.BeginDraw()
	painter := shape_drawing.NewPainter(adapter)
	painter.Draw(rect)
	adapter.EndDraw()

	assert.Equal(t, expected, buf.String())
}
