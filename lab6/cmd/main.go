package cmd

import (
	"ood/lab6/graphics"
	"ood/lab6/shape_drawing"
)

func main() {
	paintPictureOnCanvas()
}

func paintPictureOnCanvas() {
	canvas := graphics.New()
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
