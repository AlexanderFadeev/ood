package beverage

import "fmt"

type LatteSize int

const (
	LatteNormal = LatteSize(iota)
	LatteDouble
)

func (ls LatteSize) String() string {
	switch ls {
	case LatteNormal:
		return "Normal"
	case LatteDouble:
		return "Double"
	default:
		panic("Invalid latte size")
	}
}

func (ls LatteSize) GetLatteCost() float64 {
	switch ls {
	case LatteNormal:
		return 90
	case LatteDouble:
		return 130
	default:
		panic("Invalid latte size")
	}
}

type latte struct {
	coffee
}

func NewLatte(size LatteSize) Beverage {
	return &latte{*newCoffee(
		fmt.Sprintf("%s latte", size),
		90,
	)}
}
