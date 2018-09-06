package duck

import (
	"ood/lab1/duck/dance_strategy"
	"ood/lab1/duck/fly_strategy"
	"ood/lab1/duck/quack_strategy"
)

type Redhead struct {
	ConfigurableDuck
}

func NewRedheadDuck() *Redhead {
	return &Redhead{
		newDuck("redhead", new(quack_strategy.Quack), new(fly_strategy.WithWings), new(dance_strategy.Minuet)),
	}
}
