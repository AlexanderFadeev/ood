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
	syrupType SyrupType
}

func NewSyrup(syrupType SyrupType) Condiment {
	return &syrup{
		syrupType: syrupType,
	}
}

func (s *syrup) String() string {
	return fmt.Sprintf("%s syrup", s.syrupType)
}

func (syrup) GetCondimentCost() float64 {
	return 15
}
