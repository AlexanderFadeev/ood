package duck

import (
	"fmt"

	"ood/lab1/duck/dance_strategy"
	"ood/lab1/duck/fly_strategy"
	"ood/lab1/duck/quack_strategy"
)

type Duck interface {
	Quack()
	Fly()
	Dance()

	fmt.Stringer
}

type ConfigurableDuck interface {
	Duck

	SetQuackStrategy(quack_strategy.QuackStrategy)
	SetFlyStrategy(fly_strategy.FlyStrategy)
}

type duck struct {
	quack_strategy.QuackStrategy
	fly_strategy.FlyStrategy
	dance dance_strategy.DanceStrategy

	name string
}

func newDuck(name string, quack quack_strategy.QuackStrategy,
	fly fly_strategy.FlyStrategy, dance dance_strategy.DanceStrategy) ConfigurableDuck {
	return &duck{
		QuackStrategy: quack,
		FlyStrategy:   fly,
		dance:         dance,

		name: name,
	}
}

func (d *duck) SetQuackStrategy(strategy quack_strategy.QuackStrategy) {
	d.QuackStrategy = strategy
}

func (d *duck) SetFlyStrategy(strategy fly_strategy.FlyStrategy) {
	d.FlyStrategy = strategy
}

func (d *duck) SetDanceStrategy(strategy dance_strategy.DanceStrategy) {
	d.dance = strategy
}

func (d *duck) Dance() {
	d.dance()
}

func (d *duck) String() string {
	return fmt.Sprintf("I'm %s duck", d.name)
}
