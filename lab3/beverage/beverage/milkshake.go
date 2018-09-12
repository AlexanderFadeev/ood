package beverage

import "fmt"

type MilkshakeSize int

const (
	MilkshakeSmall = MilkshakeSize(iota)
	MilkshakeMedium
	MilkshakeLarge
)

func (ms MilkshakeSize) String() string {
	switch ms {
	case MilkshakeSmall:
		return "Small"
	case MilkshakeMedium:
		return "Medium"
	case MilkshakeLarge:
		return "Large"
	default:
		panic("Invalid milkshake size")
	}
}

func (ms MilkshakeSize) GetMilkshakeCost() float64 {
	switch ms {
	case MilkshakeSmall:
		return 50
	case MilkshakeMedium:
		return 60
	case MilkshakeLarge:
		return 80
	default:
		panic("Invalid milkshake size")
	}
}

type milkshake struct {
	beverage
}

func NewMilkshake(size MilkshakeSize) Beverage {
	return &milkshake{*newBeverage(
		fmt.Sprintf("%s milkshake", size),
		size.GetMilkshakeCost(),
	)}
}
