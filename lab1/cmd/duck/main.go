package main

import (
	"fmt"
	"ood/lab1/duck/dance_strategy"
	"ood/lab1/duck/quack_strategy"

	"ood/lab1/duck"
	"ood/lab1/duck/fly_strategy"
)

func main() {
	mallardDuck := duck.NewMallardDuck()
	playWithDuck(mallardDuck)

	redheadDuck := duck.NewRedheadDuck()
	playWithDuck(redheadDuck)

	rubberDuck := duck.NewRubberDuck()
	playWithDuck(rubberDuck)

	decoyDuck := duck.NewDecoyDuck()
	playWithDuck(decoyDuck)

	modelDuck := duck.NewModelDuck()
	playWithDuck(modelDuck)
	modelDuck.SetQuackStrategy(quack_strategy.Squeak)
	modelDuck.SetFlyStrategy(fly_strategy.WithWings())
	modelDuck.SetQuackStrategy(dance_strategy.Waltz)
	playWithDuck(modelDuck)
}

func playWithDuck(duck duck.Duck) {
	fmt.Println(duck)

	duck.Quack()
	duck.Dance()
	duck.Fly()

	fmt.Println()
}
