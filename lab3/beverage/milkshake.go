package beverage

type milkshake struct {
	beverage
}

func NewMilkshake() Beverage {
	return &milkshake{*newBeverage("Milkshake", 80)}
}
