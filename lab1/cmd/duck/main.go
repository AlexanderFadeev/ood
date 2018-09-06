package main

import (
	"fmt"

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
	modelDuck.SetFlyStrategy(new(fly_strategy.WithWings))
	playWithDuck(modelDuck)
}

func playWithDuck(duck duck.Duck) {
	fmt.Println(duck)

	duck.Quack()
	duck.Dance()
	duck.Fly()

	fmt.Println()
}
