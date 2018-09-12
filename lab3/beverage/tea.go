package beverage

type tea struct {
	*beverage
}

func NewTea() Beverage {
	return &tea{newBeverage("Tea", 30)}
}
