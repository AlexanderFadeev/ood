package duck

import (
	"github.com/AlexanderFadeev/ood/lab1/duck/dance_strategy"
	"github.com/AlexanderFadeev/ood/lab1/duck/fly_strategy"
	"github.com/AlexanderFadeev/ood/lab1/duck/quack_strategy"
)

type Redhead struct {
	ConfigurableDuck
}

func NewRedheadDuck() *Redhead {
	return &Redhead{
		newDuck("redhead", quack_strategy.Quack, fly_strategy.WithWings(), dance_strategy.Minuet),
	}
}
