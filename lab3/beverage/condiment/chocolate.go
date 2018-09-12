package condiment

import (
	"fmt"
)

type chocolate struct {
	condiment
}

func NewChocolate(quantity uint) Condiment {
	if quantity > 5 {
		panic(fmt.Sprintf("Invalid chocolate segments count `%d`", quantity))
	}

	return &chocolate{
		*newQuantifiedCondiment("Chocolate", 10, quantity),
	}
}
