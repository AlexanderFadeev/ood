package duck

import (
	"github.com/AlexanderFadeev/ood/lab1/duck/dance_strategy"
	"github.com/AlexanderFadeev/ood/lab1/duck/fly_strategy"
	"github.com/AlexanderFadeev/ood/lab1/duck/quack_strategy"
)

type Decoy struct {
	ConfigurableDuck
}

func NewDecoyDuck() *Decoy {
	return &Decoy{
		newDuck("decoy", quack_strategy.Muted, fly_strategy.NoWay, dance_strategy.NoWay),
	}
}
