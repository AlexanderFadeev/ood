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
		newDuck("decoy", quack_strategy.Muted, new(fly_strategy.NoWay), dance_strategy.NoWay),
	}
}

func (Decoy) Dance() {
	//Do nothing
}
