package beverage

import "fmt"

type MilkshakeSize int

const (
	MilkshakeSizeSmall = MilkshakeSize(iota)
	MilkshakeSizeMedium
	MilkshakeSizeLarge
)

func (m MilkshakeSize) String() string {
	switch m {
	case MilkshakeSizeSmall:
		return "Small"
	case MilkshakeSizeMedium:
		return "Medium"
	case MilkshakeSizeLarge:
		return "Large"
	default:
		panic("Invalid milkshake size value")
	}
}

type milkshake struct {
	size MilkshakeSize
}

func NewMilkshake(size MilkshakeSize) Beverage {
	return &milkshake{
		size: size,
	}
}

func (m *milkshake) String() string {
	return fmt.Sprintf("%s milkshake", m.size)
}

func (m *milkshake) GetCost() float64 {
	switch m.size {
	case MilkshakeSizeSmall:
		return 50
	case MilkshakeSizeMedium:
		return 60
	case MilkshakeSizeLarge:
		return 80
	default:
		panic("Invalid milkshake size value")
	}
}
