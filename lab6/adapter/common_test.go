package adapter

import (
	"ood/lab6/shape_drawing"
)

var (
	rect     = shape_drawing.NewRectangle(shape_drawing.Point{3, 14}, 5, 7)
	expected = `<draw>
	<line fromX="3" fromY="14" toX="8" toY="14"/>
	<line fromX="8" fromY="14" toX="8" toY="21"/>
	<line fromX="8" fromY="21" toX="3" toY="21"/>
	<line fromX="3" fromY="21" toX="3" toY="14"/>
</draw>
`
)
