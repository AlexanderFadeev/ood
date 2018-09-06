package quack_strategy

import "fmt"

type Squeak struct{}

func (Squeak) Quack() {
	fmt.Println("Squeak!!!")
}
