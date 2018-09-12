package condiment

import "fmt"

type SyrupType int

const (
	ChocolateSyrup = SyrupType(iota)
	MapleSyrup
)

func (st SyrupType) String() string {
	switch st {
	case ChocolateSyrup:
		return "Chocolate"
	case MapleSyrup:
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
