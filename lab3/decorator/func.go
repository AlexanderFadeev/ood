package decorator

import (
	"ood/lab3/beverage"
	"ood/lab3/condiment"
)

type DecoratorFunc func(beverage beverage.Beverage) beverage.Beverage

func MakeDecoratorFunc(condiment condiment.Condiment) DecoratorFunc {
	return func(beverage beverage.Beverage) beverage.Beverage {
		return Decorate(beverage, condiment)
	}
}
