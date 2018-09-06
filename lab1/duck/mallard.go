package duck

import (
	"ood/lab1/duck/dance_strategy"
	"ood/lab1/duck/fly_strategy"
	"ood/lab1/duck/quack_strategy"
)

type Mallard struct {
	ConfigurableDuck
}

func NewMallardDuck() *Mallard {
	return &Mallard{
		newDuck("mallard", quack_strategy.Quack, new(fly_strategy.WithWings), dance_strategy.Waltz),
	}
}
