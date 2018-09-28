package canvas

import (
	"context"
	"image"
	"net/http"
	"sync"

	"ood/lab4/color"
	"ood/lab4/point"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/image/bmp"
)

const pixelRadius = 0.5

const (
	address     = ":8080"
	fullAddress = "http://localhost:8080/"
)

type Canvas interface {
	SetColor(color.Color)
	DrawLine(from, to point.Point)
	DrawEllipse(leftTop point.Point, width, height float64)
}

type DisplayableCanvas interface {
	Canvas
	Display()
}

type pointCanvas interface {
	drawPoint(point.Point)
}

type canvas struct {
	color color.Color
	img   *image.RGBA
}

func New(width, height uint) DisplayableCanvas {
	rect := image.Rect(0, 0, int(width), int(height))
	return &canvas{
		img: image.NewRGBA(rect),
	}
}

func (c *canvas) SetColor(color color.Color) {
	c.color = color
}

func (c *canvas) DrawLine(from, to point.Point) {
	line := line{
		from: from,
		to:   to,
	}
	line.draw(c)
}

func (c *canvas) DrawEllipse(leftTop point.Point, width, height float64) {
	center := point.Point{
		X: leftTop.X + width/2,
		Y: leftTop.Y + height/2,
	}
	ellipse := ellipse{
		center:           center,
		horizontalRadius: width / 2,
		verticalRadius:   height / 2,
	}
	ellipse.draw(c)
}

func (pc *canvas) drawPoint(p point.Point) {
	pc.img.Set(int(p.X), int(p.Y), pc.color)
}

func (c *canvas) Display() {
	router := mux.NewRouter()

	var wg sync.WaitGroup
	wg.Add(1)

	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "image/bmp")
		bmp.Encode(w, c.img)
		wg.Done()
	}))

	server := http.Server{
		Addr:    address,
		Handler: router,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logrus.Error(err)
		}
	}()
	defer server.Shutdown(context.Background())

	err := open.Start(fullAddress)
	if err != nil {
		logrus.Error(err)
		return
	}

	wg.Wait()
}
