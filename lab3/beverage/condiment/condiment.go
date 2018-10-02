package condiment

import (
	"fmt"
)

type Condiment interface {
	fmt.Stringer

	GetCondimentCost() float64
}
