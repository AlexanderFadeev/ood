package beverage

type latte struct {
	*coffee
}

func NewLatte() Beverage {
	return &latte{newCoffee("Latte", 90)}
}
