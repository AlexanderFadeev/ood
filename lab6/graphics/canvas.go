package graphics

import "fmt"

type Canvas interface {
	SetColor(rgbColor uint32)
	MoveTo(x, y int)
	LineTo(x, y int)
}

type canvas struct{}

func New() Canvas {
	return new(canvas)
}

func (canvas) SetColor(rgb uint32) {
	fmt.Printf("SetColor(#%06X)\n", rgb)
}

func (canvas) MoveTo(x, y int) {
	fmt.Printf("MoveTo(%d, %d)\n", x, y)
}

func (canvas) LineTo(x, y int) {
	fmt.Printf("LineTo(%d, %d)\n", x, y)
}
