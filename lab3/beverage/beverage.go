package beverage

import "fmt"

type Beverage interface {
	fmt.Stringer

	GetCost() float64
}

type beverage struct {
	description string
	cost        float64
}

func newBeverage(description string, cost float64) *beverage {
	return &beverage{
		description: description,
		cost:        cost,
	}
}

func (b *beverage) String() string {
	return b.description
}

func (b *beverage) GetCost() float64 {
	return b.cost
}
