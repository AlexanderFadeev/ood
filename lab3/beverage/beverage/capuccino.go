package beverage

import "fmt"

type CapuccinoSize int

const (
	CapuccinoNormal = CapuccinoSize(iota)
	CapuccinoDouble
)

func (ls CapuccinoSize) String() string {
	switch ls {
	case CapuccinoNormal:
		return "Normal"
	case CapuccinoDouble:
		return "Double"
	default:
		panic("Invalid capuccino size")
	}
}

func (ls CapuccinoSize) GetCapuccinoCost() float64 {
	switch ls {
	case CapuccinoNormal:
		return 80
	case CapuccinoDouble:
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
		80,
	)}
}
