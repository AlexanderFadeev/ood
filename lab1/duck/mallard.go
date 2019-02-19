package duck

import (
	"github.com/AlexanderFadeev/ood/lab1/duck/dance_strategy"
	"github.com/AlexanderFadeev/ood/lab1/duck/fly_strategy"
	"github.com/AlexanderFadeev/ood/lab1/duck/quack_strategy"
)

type Mallard struct {
	ConfigurableDuck
}

func NewMallardDuck() *Mallard {
	return &Mallard{
		newDuck("mallard", quack_strategy.Quack, fly_strategy.WithWings(), dance_strategy.Waltz),
	}
}
