package adapter

import (
	"fmt"

	"github.com/AlexanderFadeev/ood/lab6/shape_drawing"
)

var (
	rect          = shape_drawing.NewRectangle(shape_drawing.Point{3, 14}, 5, 7, 0x425262)
	expectedColor = `<color r="0.26" g="0.32" b="0.38" a="1.00" />`
	expected      = fmt.Sprintf(`<draw>
	<line fromX="3" fromY="14" toX="8" toY="14">
		%s
	</line>
	<line fromX="8" fromY="14" toX="8" toY="21">
		%s
	</line>
	<line fromX="8" fromY="21" toX="3" toY="21">
		%s
	</line>
	<line fromX="3" fromY="21" toX="3" toY="14">
		%s
	</line>
</draw>
`, expectedColor, expectedColor, expectedColor, expectedColor)
)
