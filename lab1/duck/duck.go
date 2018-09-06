package duck

import (
	"fmt"

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

	name string
}

func newDuck(name string,
	quackStrategy quack_strategy.QuackStrategy, flyStrategy fly_strategy.FlyStrategy) ConfigurableDuck {
	return &duck{
		QuackStrategy: quackStrategy,
		FlyStrategy:   flyStrategy,

		name: name,
	}
}

func (d *duck) SetQuackStrategy(strategy quack_strategy.QuackStrategy) {
	d.QuackStrategy = strategy
}

func (d *duck) SetFlyStrategy(strategy fly_strategy.FlyStrategy) {
	d.FlyStrategy = strategy
}

func (duck) Dance() {
	fmt.Println("I'm dancing!!!")
}

func (d *duck) String() string {
	return fmt.Sprintf("I'm %s duck", d.name)
}
