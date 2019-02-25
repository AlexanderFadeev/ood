package main

import (
	"github.com/AlexanderFadeev/ood/lab4/canvas"
	"github.com/AlexanderFadeev/ood/lab4/color"
	"github.com/AlexanderFadeev/ood/lab4/point"
	"github.com/AlexanderFadeev/ood/lab4/shape"
	"github.com/sirupsen/logrus"
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
