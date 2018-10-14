package main

import (
	"bufio"
	"fmt"
	"os"

	"ood/lab6/adapter"
	"ood/lab6/graphics"
	"ood/lab6/modern_graphics"
	"ood/lab6/shape_drawing"
)

func main() {
	if promtYesNo("Should we use new API?") {
		fmt.Println("Using modern graphics lib")
		paintPictureOnModernGraphicsRenderer()
	} else {
		fmt.Println("Using old graphics lib")
		paintPictureOnCanvas()
	}
}

func promtYesNo(question string) bool {
	fmt.Print(question + " (y/n) ")
	stdinReader := bufio.NewReader(os.Stdin)
	ch, _, _ := stdinReader.ReadRune()
	return ch == 'y' || ch == 'Y'
}

func paintPictureOnCanvas() {
	canvas := graphics.New()
	painter := shape_drawing.NewPainter(canvas)
	paintPicture(painter)
}

func paintPictureOnModernGraphicsRenderer() {
	renderer := modern_graphics.NewRenderer(os.Stdout)
	renderer.BeginDraw()
	defer renderer.EndDraw()

	canvas := adapter.NewObjectAdapter(renderer)
	painter := shape_drawing.NewPainter(canvas)

	paintPicture(painter)
}

func paintPicture(painter shape_drawing.Painter) {
	triangle := shape_drawing.NewTriangle(
		shape_drawing.Point{10, 15},
		shape_drawing.Point{100, 200},
		shape_drawing.Point{150, 250},
	)
	rect := shape_drawing.NewRectangle(shape_drawing.Point{30, 40}, 18, 24)

	painter.Draw(triangle)
	painter.Draw(rect)
}
