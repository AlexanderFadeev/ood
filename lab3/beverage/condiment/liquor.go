package condiment

import "fmt"

type LiquorType int

const (
	LiquorTypeNut = LiquorType(iota)
	LiquorTypeChocolate
)

func (lt LiquorType) String() string {
	switch lt {
	case LiquorTypeNut:
		return "Nut"
	case LiquorTypeChocolate:
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
