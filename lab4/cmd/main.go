package main

import (
	"github.com/AlexanderFadeev/ood/lab4/canvas"
	"github.com/AlexanderFadeev/ood/lab4/designer"
	"github.com/AlexanderFadeev/ood/lab4/painter"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	width  = 1366
	height = 768

	filename = "shapes.dat"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	c := canvas.New(width, height)

	file, err := os.Open(filename)
	if err != nil {
		logrus.Fatal(err)
	}
	defer file.Close()

	d := designer.New()
	draft, err := d.CreateDraft(file)
	if err != nil {
		logrus.Fatal(err)
	}

	p := painter.New()
	p.Paint(draft, c)

	c.Display()
}
