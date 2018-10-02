package beverage

import "fmt"

type Beverage interface {
	fmt.Stringer

	GetCost() float64
}
