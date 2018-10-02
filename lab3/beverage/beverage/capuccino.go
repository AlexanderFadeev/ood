package beverage

import "fmt"

type CapuccinoSize int

const (
	CapuccinoSizeRegular = CapuccinoSize(iota)
	CapuccinoSizeDouble
)

func (cs CapuccinoSize) String() string {
	switch cs {
	case CapuccinoSizeRegular:
		return "Regular"
	case CapuccinoSizeDouble:
		return "Double"
	default:
		panic("Invalid capuccino size value")
	}
}

type capuccino struct {
	size CapuccinoSize
}

func NewCapuccino(size CapuccinoSize) Beverage {
	return &capuccino{
		size: size,
	}
}

func (c *capuccino) String() string {
	return fmt.Sprintf("%s capuccino", c.size)
}

func (c *capuccino) GetCost() float64 {
	switch c.size {
	case CapuccinoSizeRegular:
		return 80
	case CapuccinoSizeDouble:
		return 120
	default:
		panic("Invalid capuccino size value")
	}
}
