package adapter

import "image/color"

func rgbToRGBA(rgb uint32) color.Color {
	r := uint8((rgb & 0xFF0000) >> 16)
	g := uint8((rgb & 0x00FF00) >> 8)
	b := uint8((rgb & 0x0000FF) >> 0)

	return &color.RGBA{r, g, b, 0xFF}
}
