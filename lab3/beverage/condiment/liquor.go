package condiment

import "fmt"

type LiquorType int

const (
	LiquorNut = LiquorType(iota)
	LiquorChocolate
)

func (lt LiquorType) String() string {
	switch lt {
	case LiquorNut:
		return "Nut"
	case LiquorChocolate:
		return "Chocolate"
	default:
		panic("Invalid liquor type")
	}
}

type liquor struct {
	condiment
}

func NewLiquor(liquorType LiquorType) Condiment {
	return &liquor{*newCondiment(
		fmt.Sprintf("%s liquor", liquorType),
		50,
	)}
}
