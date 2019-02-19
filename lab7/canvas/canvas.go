package canvas

import (
	"fmt"
	"image/color"
	"io"

	"github.com/AlexanderFadeev/ood/lab7/point"
	"github.com/AlexanderFadeev/ood/lab7/rect"

	"github.com/AlexanderFadeev/colors"
	"github.com/ajstarks/svgo"
)

type Canvas interface {
	SetLineColor(color.Color)
	SetFillColor(color.Color)
	SetLineWidth(float64)

	MoveTo(point.Point)
	LineTo(point.Point)

	DrawEllipse(rect.Rect)

	End()
}

type canvas struct {
	impl      *svg.SVG
	polyLineX []int
	polyLineY []int
	lineColor color.Color
	fillColor color.Color
	lineWidth float64
}

func New(writer io.Writer, w, h int) Canvas {
	impl := svg.New(writer)
	impl.Start(w, h)
	return &canvas{
		impl: impl,
	}
}

func (c *canvas) SetLineColor(color color.Color) {
	c.finishPolyLine()
	c.lineColor = color
}

func (c *canvas) SetFillColor(color color.Color) {
	c.finishPolyLine()
	c.fillColor = color
}

func (c *canvas) SetLineWidth(lineWidth float64) {
	c.finishPolyLine()
	c.lineWidth = lineWidth
}

func (c *canvas) MoveTo(p point.Point) {
	c.finishPolyLine()
	c.addPolyLineVertex(p)
}

func (c *canvas) LineTo(p point.Point) {
	c.addPolyLineVertex(p)
}

func (c *canvas) addPolyLineVertex(p point.Point) {
	c.polyLineX = append(c.polyLineX, int(p.X))
	c.polyLineY = append(c.polyLineY, int(p.Y))
}

func (c *canvas) DrawEllipse(rect rect.Rect) {
	c.finishPolyLine()
	c.impl.Ellipse(
		int(rect.LeftTop.X), int(rect.LeftTop.Y),
		int(rect.Width()), int(rect.Height()),
		c.getStyleStr(),
	)
}

func (c *canvas) End() {
	c.finishPolyLine()
	c.impl.End()
}

func (c *canvas) finishPolyLine() {
	if len(c.polyLineX) == 0 {
		return
	}

	c.impl.Polyline(c.polyLineX, c.polyLineY, c.getStyleStr())

	c.polyLineX = []int{}
	c.polyLineY = []int{}
}

func (c *canvas) getStyleStr() string {
	return fmt.Sprintf("stroke:%s;fill:%s;stroke-width:%.2f;stroke-linecap:round;stroke-linejoin:round",
		c.formatColor(c.lineColor), c.formatColor(c.fillColor), c.lineWidth)
}

func (canvas) formatColor(color color.Color) string {
	if color == nil {
		return "none"
	}
	return colors.FromStdColor(color).ToRGB().String()
}
