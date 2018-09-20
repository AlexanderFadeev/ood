package beverage

import "fmt"

type TeaType int

const (
	TeaTypeBlack = TeaType(iota)
	TeaTypeGreen
	TeaTypeFruit
	TeaTypeRed
)

func (tt TeaType) String() string {
	switch tt {
	case TeaTypeBlack:
		return "Black"
	case TeaTypeGreen:
		return "Green"
	case TeaTypeFruit:
		return "Fruit"
	case TeaTypeRed:
		return "Red"
	default:
		panic("Invalid tea type")
	}
}

type tea struct {
	beverage
}

func NewTea(teaType TeaType) Beverage {
	return &tea{*newBeverage(
		fmt.Sprintf("%s tea", teaType),
		30,
	)}
}
