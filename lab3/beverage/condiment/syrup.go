package condiment

import "fmt"

type SyrupType int

const (
	SyrupTypeChocolate = SyrupType(iota)
	SyrupTypeMaple
)

func (st SyrupType) String() string {
	switch st {
	case SyrupTypeChocolate:
		return "Chocolate"
	case SyrupTypeMaple:
		return "Maple"
	default:
		panic("Invalid syrup type")
	}
}

type syrup struct {
	condiment
}

func NewSyrup(syrupType SyrupType) Condiment {
	return &syrup{*newCondiment(
		fmt.Sprintf("%s syrup", syrupType),
		15,
	)}
}
