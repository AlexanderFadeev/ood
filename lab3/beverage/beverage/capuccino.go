package beverage

import "fmt"

type CapuccinoSize int

const (
	CapuccinoSizeNormal = CapuccinoSize(iota)
	CapuccinoSizeDouble
)

func (ls CapuccinoSize) String() string {
	switch ls {
	case CapuccinoSizeNormal:
		return "Normal"
	case CapuccinoSizeDouble:
		return "Double"
	default:
		panic("Invalid capuccino size")
	}
}

func (ls CapuccinoSize) GetCapuccinoCost() float64 {
	switch ls {
	case CapuccinoSizeNormal:
		return 80
	case CapuccinoSizeDouble:
		return 120
	default:
		panic("Invalid capuccino size")
	}
}

type capuccino struct {
	coffee
}

func NewCapuccino(size CapuccinoSize) Beverage {
	return &capuccino{*newCoffee(
		fmt.Sprintf("%s capuccino", size),
		size.GetCapuccinoCost(),
	)}
}
