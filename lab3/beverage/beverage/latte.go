package beverage

import "fmt"

type LatteSize int

const (
	LatteSizeRegular = LatteSize(iota)
	LatteSizeDouble
)

func (ls LatteSize) String() string {
	switch ls {
	case LatteSizeRegular:
		return "Regular"
	case LatteSizeDouble:
		return "Double"
	default:
		panic("Invalid latte size value")
	}
}

type latte struct {
	size LatteSize
}

func NewLatte(size LatteSize) Beverage {
	return &latte{
		size: size,
	}
}

func (l *latte) String() string {
	return fmt.Sprintf("%s latte", l.size)
}

func (l *latte) GetCost() float64 {
	switch l.size {
	case LatteSizeRegular:
		return 90
	case LatteSizeDouble:
		return 130
	default:
		panic("Invalid latte size value")
	}
}
