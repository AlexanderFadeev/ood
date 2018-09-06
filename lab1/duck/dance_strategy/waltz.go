package dance_strategy

import "fmt"

type Waltz struct{}

func (Waltz) Dance() {
	fmt.Println("I'm dancing waltz")
}
