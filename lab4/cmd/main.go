package main

import (
	"github.com/sirupsen/logrus"
	"ood/lab4/canvas"
	"ood/lab4/color"
	"ood/lab4/point"
	"ood/lab4/shape"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	c := canvas.New(1366, 768)
	c.SetColor(color.Red)
	c.DrawLine(point.Point{0, 0}, point.Point{1366, 768})
	c.SetColor(color.Green)
	c.DrawEllipse(point.Point{100, 100}, 500, 100)

	for i := uint(3); i < 64; i++ {
		regPol, _ := shape.NewRegularPolygon(i, point.Point{700, 350}, 350, color.Blue)
		regPol.Draw(c)
	}

	c.Display()
}
