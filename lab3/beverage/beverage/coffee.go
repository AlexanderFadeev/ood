package beverage

type coffee struct {
	beverage
}

func newCoffee(name string, cost float64) *coffee {
	return &coffee{*newBeverage(name, cost)}
}

func NewCoffee() Beverage {
	return newCoffee("Coffee", 60)
}
