package modern_graphics

import (
	"fmt"
	"image/color"
	"io"
)

type Renderer interface {
	BeginDraw()
	DrawLine(from, to Point, color color.Color)
	EndDraw()
}

type renderer struct {
	writer  io.Writer
	drawing bool
}

func NewRenderer(writer io.Writer) Renderer {
	return &renderer{
		writer:  writer,
		drawing: false,
	}
}

func (r *renderer) BeginDraw() {
	if r.drawing {
		panic("Drawing is already begun")
	}
	io.WriteString(r.writer, "<draw>\n")
	r.drawing = true
}

func (r *renderer) DrawLine(from, to Point, color color.Color) {
	if !r.drawing {
		panic("DrawLine is allowed between BeginDraw()/EndDraw() only")
	}

	rc, g, b, a := color.RGBA()
	io.WriteString(r.writer, fmt.Sprintf(
		`	<line fromX="%d" fromY="%d" toX="%d" toY="%d">
		<color r="%.2f" g="%.2f" b="%.2f" a="%.2f" />
	</line>
`,
		from.X, from.Y, to.X, to.Y, float32(rc)/(1<<16), float32(g)/(1<<16), float32(b)/(1<<16), float32(a)/(1<<16),
	))
}

func (r *renderer) EndDraw() {
	if !r.drawing {
		panic("Drawing has not been started")
	}
	io.WriteString(r.writer, "</draw>\n")
	r.drawing = false
}
