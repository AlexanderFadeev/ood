package beverage

import "fmt"

type TeaKind int

const (
	TeaKindBlack = TeaKind(iota)
	TeaKindGreen
	TeaKindFruit
	TeaKindRed
)

func (t TeaKind) String() string {
	switch t {
	case TeaKindBlack:
		return "Black"
	case TeaKindGreen:
		return "Green"
	case TeaKindFruit:
		return "Fruit"
	case TeaKindRed:
		return "Red"
	default:
		panic("Invalid tea kind value")
	}
}

type tea struct {
	kind TeaKind
}

func NewTea(kind TeaKind) Beverage {
	return &tea{
		kind: kind,
	}
}

func (t *tea) String() string {
	return fmt.Sprintf("%s tea", t.kind)
}

func (tea) GetCost() float64 {
	return 30
}
