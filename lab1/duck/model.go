package duck

import (
	"ood/lab1/duck/fly_strategy"
	"ood/lab1/duck/quack_strategy"
)

type Model struct {
	ConfigurableDuck
}

func NewModelDuck() *Model {
	return &Model{
		newDuck("model", new(quack_strategy.Quack), new(fly_strategy.NoWay)),
	}
}

func (Model) Dance() {
	//Do nothing
}
