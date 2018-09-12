package beverage

import "fmt"

type TeaType int

const (
	TeaBlack = TeaType(iota)
	TeaGreen
	TeaFruit
	TeaRed
)

func (tt TeaType) String() string {
	switch tt {
	case TeaBlack:
		return "Black"
	case TeaGreen:
		return "Green"
	case TeaFruit:
		return "Fruit"
	case TeaRed:
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
