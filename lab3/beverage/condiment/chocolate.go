package condiment

import (
	"fmt"
)

type chocolate struct {
	quantity uint
}

func NewChocolate(quantity uint) Condiment {
	if !(0 < quantity && quantity <= 5) {
		panic(fmt.Sprintf("Invalid chocolate segments count `%d`", quantity))
	}

	return &chocolate{
		quantity: quantity,
	}
}

func (c *chocolate) String() string {
	return fmt.Sprintf(quantifiedDescriptionFormat, "Chocolate", c.quantity)
}

func (c *chocolate) GetCondimentCost() float64 {
	return 10 * float64(c.quantity)
}
