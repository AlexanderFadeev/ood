package main

import (
	"fmt"
	"os"

	"ood/cli_util"
	"ood/lab6/adapter"
	"ood/lab6/graphics"
	"ood/lab6/modern_graphics"
	"ood/lab6/shape_drawing"
)

func main() {
	if !cli_util.PromtYesNo("Should we use new API?") {
		fmt.Println("Using old graphics lib")
		paintPictureOnCanvas()
	} else if !cli_util.PromtYesNo("Should we use class adapter?") {
		fmt.Println("Using modern graphics lib with object adapter")
		paintPictureOnModernGraphicsRendererObjectAdapter()
	} else {
		fmt.Println("Using modern graphics lib with class adapter")
		paintPictureOnModernGraphicsRendererClassAdapter()
	}
}

func paintPictureOnCanvas() {
	canvas := graphics.New()
	painter := shape_drawing.NewPainter(canvas)
	paintPicture(painter)
}

func paintPictureOnModernGraphicsRendererObjectAdapter() {
	renderer := modern_graphics.NewRenderer(os.Stdout)
	renderer.BeginDraw()
	defer renderer.EndDraw()

	canvas := adapter.NewCanvasObjectAdapter(renderer)
	painter := shape_drawing.NewPainter(canvas)

	paintPicture(painter)
}

func paintPictureOnModernGraphicsRendererClassAdapter() {
	classAdapter := adapter.NewCanvasClassAdapter(os.Stdout)
	classAdapter.BeginDraw()
	defer classAdapter.EndDraw()

	painter := shape_drawing.NewPainter(classAdapter)

	paintPicture(painter)
}

func paintPicture(painter shape_drawing.Painter) {
	triangle := shape_drawing.NewTriangle(
		shape_drawing.Point{10, 15},
		shape_drawing.Point{100, 200},
		shape_drawing.Point{150, 250},
		0xFF0000,
	)
	rect := shape_drawing.NewRectangle(shape_drawing.Point{30, 40}, 18, 24, 0x00FF00)

	painter.Draw(triangle)
	painter.Draw(rect)
}
