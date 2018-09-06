package duck

import (
	"ood/lab1/duck/dance_strategy"
	"ood/lab1/duck/fly_strategy"
	"ood/lab1/duck/quack_strategy"
)

type Rubber struct {
	ConfigurableDuck
}

func NewRubberDuck() *Rubber {
	return &Rubber{
		newDuck("rubber", quack_strategy.Squeak, fly_strategy.NoWay, dance_strategy.NoWay),
	}
}

func (Rubber) Dance() {
	//Do nothing
}
