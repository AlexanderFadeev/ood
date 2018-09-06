package quack_strategy

import "fmt"

type Quack struct{}

func (Quack) Quack() {
	fmt.Println("Quack Quack!!!")
}
