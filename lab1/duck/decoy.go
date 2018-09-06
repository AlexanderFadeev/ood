package duck

import (
	"ood/lab1/duck/dance_strategy"
	"ood/lab1/duck/fly_strategy"
	"ood/lab1/duck/quack_strategy"
)

type Decoy struct {
	ConfigurableDuck
}

func NewDecoyDuck() *Decoy {
	return &Decoy{
		newDuck("decoy", new(quack_strategy.Muted), new(fly_strategy.NoWay), new(dance_strategy.NoWay)),
	}
}

func (Decoy) Dance() {
	//Do nothing
}
