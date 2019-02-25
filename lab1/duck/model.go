package duck

import (
	"github.com/AlexanderFadeev/ood/lab1/duck/dance_strategy"
	"github.com/AlexanderFadeev/ood/lab1/duck/fly_strategy"
	"github.com/AlexanderFadeev/ood/lab1/duck/quack_strategy"
)

type Model struct {
	ConfigurableDuck
}

func NewModelDuck() *Model {
	return &Model{
		newDuck("model", quack_strategy.Quack, fly_strategy.NoWay, dance_strategy.NoWay),
	}
}
