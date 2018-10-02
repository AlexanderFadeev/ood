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
	liquorType LiquorType
}

func NewLiquor(liquorType LiquorType) Condiment {
	return &liquor{
		liquorType: liquorType,
	}
}

func (l *liquor) String() string {
	return fmt.Sprintf("%s liquor", l.liquorType)
}

func (liquor) GetCondimentCost() float64 {
	return 50
}
