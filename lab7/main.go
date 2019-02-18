package main

import (
	"io"

	"ood/browser_display"
	"ood/lab7/canvas"
	"ood/lab7/point"
	"ood/lab7/rect"
	"ood/lab7/shape"
	"ood/lab7/slide"

	"golang.org/x/image/colornames"
)

func main() {
	r := getSlideRect()
	s := slide.NewSlide(r)

	s.SetBackgroundColor(colornames.Deepskyblue)

	s.InsertShape(getGrassShape(), 0)
	s.InsertShape(getHouseShape(), 1)
	s.InsertShape(getSunShape(), 2)
	s.InsertShape(getPuddleShape(), 3)

	s.GetShape(1).SetFrame(rect.New(point.Point{900, 200}, point.Point{1300, 700}))

	browser_display.DisplayInBrowser(func(w io.Writer) {
		c := canvas.New(w, int(r.Width()), int(r.Height()))
		defer c.End()

		s.Draw(c)
	}, "image/svg+xml")
}

func getSlideRect() rect.Rect {
	return rect.New(
		point.Point{X: 0, Y: 0},
		point.Point{X: 1366, Y: 768},
	)
}

func getHouseShape() shape.Shape {
	house := shape.NewGroup()

	roof := shape.NewTriangle([3]point.Point{{100, 145}, {145, 100}, {190, 145}})
	roof.GetFillStyle().SetColor(colornames.Burlywood)
	roof.GetLineStyle().SetColor(colornames.Black)

	base := shape.NewRectangle(rect.New(point.Point{100, 145}, point.Point{190, 235}))
	base.GetFillStyle().SetColor(colornames.Peru)
	base.GetLineStyle().SetColor(colornames.Black)

	window := shape.NewRectangle(rect.New(point.Point{130, 175}, point.Point{160, 205}))
	window.GetFillStyle().SetColor(colornames.Lightblue)
	window.GetLineStyle().SetColor(colornames.Sandybrown)
	window.GetLineStyle().SetWidth(5)

	house.InsertShape(base, 0)
	house.InsertShape(roof, 1)
	house.InsertShape(window, 2)

	if house.InsertShape(house, 0) {
		panic("Should not be able to insert group into itself")
	}

	return house
}

func getGrassShape() shape.Shape {
	grass := shape.NewRectangle(rect.New(point.Point{-5, 768 / 2}, point.Point{1366, 773}))
	grass.GetFillStyle().SetColor(colornames.Forestgreen)
	grass.GetLineStyle().SetColor(colornames.Greenyellow)
	grass.GetLineStyle().SetWidth(5)
	return grass
}

func getSunShape() shape.Shape {
	sun := shape.NewEllipse(rect.New(point.Point{300, 100}, point.Point{350, 150}))
	sun.GetFillStyle().SetColor(colornames.Yellow)
	return sun
}

func getPuddleShape() shape.Shape {
	puddle := shape.NewEllipse(rect.New(point.Point{400, 600}, point.Point{600, 650}))
	puddle.GetFillStyle().SetColor(colornames.Cyan)
	return puddle
}
