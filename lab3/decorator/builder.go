package decorator

import (
	"ood/lab3/beverage"
	"ood/lab3/condiment"
)

type Builder interface {
	WithCondiment(condiment.Condiment) Builder
	Build() beverage.Beverage
}

type builder struct {
	beverage beverage.Beverage
}

func NewBuilder(beverage beverage.Beverage) Builder {
	return &builder{
		beverage: beverage,
	}
}

func (b *builder) WithCondiment(condiment condiment.Condiment) Builder {
	return &builder{
		beverage: Decorate(b.beverage, condiment),
	}
}

func (b *builder) Build() beverage.Beverage {
	return b.beverage
}
