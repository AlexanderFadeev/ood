package beverage

type coffee struct{}

func NewCoffee() Beverage {
	return new(coffee)
}

func (coffee) String() string {
	return "Coffee"
}

func (coffee) GetCost() float64 {
	return 60
}
