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
	SetDanceStrategy(dance_strategy.DanceStrategy)
}

type duck struct {
	quack quack_strategy.QuackStrategy
	fly   fly_strategy.FlyStrategy
	dance dance_strategy.DanceStrategy

	name string
}

func newDuck(name string, quack quack_strategy.QuackStrategy,
	fly fly_strategy.FlyStrategy, dance dance_strategy.DanceStrategy) ConfigurableDuck {
	return &duck{
		quack: quack,
		fly:   fly,
		dance: dance,

		name: name,
	}
}

func (d *duck) SetQuackStrategy(strategy quack_strategy.QuackStrategy) {
	d.quack = strategy
}

func (d *duck) SetFlyStrategy(strategy fly_strategy.FlyStrategy) {
	d.fly = strategy
}

func (d *duck) SetDanceStrategy(strategy dance_strategy.DanceStrategy) {
	d.dance = strategy
}

func (d *duck) Quack() {
	d.quack()
}

func (d *duck) Fly() {
	d.fly()
}

func (d *duck) Dance() {
	d.dance()
}

func (d *duck) String() string {
	return fmt.Sprintf("I'm %s duck", d.name)
}
