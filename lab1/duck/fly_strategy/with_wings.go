package fly_strategy

import "fmt"

type WithWings struct{}

func (WithWings) Fly() {
	fmt.Println("I'm flying with wings!!!")
}
