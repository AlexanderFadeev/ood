package duck

import (
	"ood/lab1/duck/fly_strategy"
	"ood/lab1/duck/quack_strategy"
)

type Rubber struct {
	ConfigurableDuck
}

func NewRubberDuck() *Rubber {
	return &Rubber{
		newDuck("rubber", new(quack_strategy.Squeak), new(fly_strategy.NoWay)),
	}
}

func (Rubber) Dance() {
	//Do nothing
}
