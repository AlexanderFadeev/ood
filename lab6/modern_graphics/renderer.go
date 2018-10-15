package modern_graphics

import (
	"fmt"
	"io"
)

type Renderer interface {
	BeginDraw()
	DrawLine(from, to Point)
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

func (r *renderer) DrawLine(from, to Point) {
	if !r.drawing {
		panic("DrawLine is allowed between BeginDraw()/EndDraw() only")
	}
	io.WriteString(r.writer, fmt.Sprintf(
		`	<line fromX="%d" fromY="%d" toX="%d" toY="%d"/>
`,
		from.X, from.Y, to.X, to.Y,
	))
}

func (r *renderer) EndDraw() {
	if !r.drawing {
		panic("Drawing has not been started")
	}
	io.WriteString(r.writer, "</draw>\n")
	r.drawing = false
}
