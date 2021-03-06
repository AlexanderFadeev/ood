package main

import (
	"fmt"

	"github.com/AlexanderFadeev/ood/lab1/duck"
	"github.com/AlexanderFadeev/ood/lab1/duck/dance_strategy"
	"github.com/AlexanderFadeev/ood/lab1/duck/fly_strategy"
	"github.com/AlexanderFadeev/ood/lab1/duck/quack_strategy"
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
	modelDuck.SetDanceStrategy(dance_strategy.Waltz)
	playWithDuck(modelDuck)
}

func playWithDuck(duck duck.Duck) {
	fmt.Println(duck)

	duck.Quack()
	duck.Dance()
	duck.Fly()

	fmt.Println()
}
