package decorator

import (
	"strings"

	"ood/lab3/beverage/beverage"
	"ood/lab3/beverage/condiment"
)

type condimentDecorator struct {
	beverage  beverage.Beverage
	condiment condiment.Condiment
}

func Decorate(beverage beverage.Beverage, condiment condiment.Condiment) beverage.Beverage {
	return &condimentDecorator{
		beverage:  beverage,
		condiment: condiment,
	}
}

func (cd *condimentDecorator) GetCost() float64 {
	return cd.beverage.GetCost() + cd.condiment.GetCondimentCost()
}

func (cd *condimentDecorator) String() string {
	var builder strings.Builder
	builder.WriteString(cd.beverage.String())
	builder.WriteString(", ")
	builder.WriteString(cd.condiment.String())
	return builder.String()
}
