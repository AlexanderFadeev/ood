package condiment

import "fmt"

type lemon struct {
	quantity uint
}

func NewLemon(quantity uint) Condiment {
	return &lemon{
		quantity: quantity,
	}
}

func (l *lemon) String() string {
	return fmt.Sprintf(quantifiedDescriptionFormat, "Lemon", l.quantity)
}

func (l *lemon) GetCondimentCost() float64 {
	return 10 * float64(l.quantity)
}
