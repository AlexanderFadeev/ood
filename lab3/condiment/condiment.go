package condiment

import (
	"fmt"
)

type Condiment interface {
	fmt.Stringer

	GetCondimentCost() float64
}

type condiment struct {
	description string
	cost        float64
}

func newCondiment(description string, cost float64) *condiment {
	return &condiment{
		description: description,
		cost:        cost,
	}
}

func (c *condiment) String() string {
	return c.description
}

func (c *condiment) GetCondimentCost() float64 {
	return c.cost
}

func newQuantifiedCondiment(description string, cost float64, quantity uint) *condiment {
	return newCondiment(
		fmt.Sprintf("%s x%d", description, quantity),
		cost*float64(quantity),
	)
}

func newWeightedCondiment(description string, cost float64, weight float64) *condiment {
	return newCondiment(
		fmt.Sprintf("%s %.2fg", description, weight),
		cost*weight,
	)
}
