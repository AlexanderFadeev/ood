package beverage

import "fmt"

type LatteSize int

const (
	LatteSizeNormal = LatteSize(iota)
	LatteSizeDouble
)

func (ls LatteSize) String() string {
	switch ls {
	case LatteSizeNormal:
		return "Normal"
	case LatteSizeDouble:
		return "Double"
	default:
		panic("Invalid latte size")
	}
}

func (ls LatteSize) GetLatteCost() float64 {
	switch ls {
	case LatteSizeNormal:
		return 90
	case LatteSizeDouble:
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
