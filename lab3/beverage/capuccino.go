package beverage

type capuccino struct {
	*coffee
}

func NewCapuccino() Beverage {
	return &capuccino{newCoffee("Capuccino", 80)}
}
