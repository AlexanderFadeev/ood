package beverage

import "fmt"

type MilkshakeSize int

const (
	MilkshakeSizeSmall = MilkshakeSize(iota)
	MilkshakeSizeMedium
	MilkshakeSizeLarge
)

func (ms MilkshakeSize) String() string {
	switch ms {
	case MilkshakeSizeSmall:
		return "Small"
	case MilkshakeSizeMedium:
		return "Medium"
	case MilkshakeSizeLarge:
		return "Large"
	default:
		panic("Invalid milkshake size")
	}
}

func (ms MilkshakeSize) GetMilkshakeCost() float64 {
	switch ms {
	case MilkshakeSizeSmall:
		return 50
	case MilkshakeSizeMedium:
		return 60
	case MilkshakeSizeLarge:
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
