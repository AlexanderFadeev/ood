package condiment

import "fmt"

type IceCubeType int

const (
	IceCubeTypeDry = IceCubeType(iota)
	IceCubeTypeWater
)

func (ict IceCubeType) String() string {
	switch ict {
	case IceCubeTypeDry:
		return "Dry"
	case IceCubeTypeWater:
		return "Water"
	default:
		panic("Invalid ice cube type")
	}
}

func (ict IceCubeType) GetCubeCost() float64 {
	switch ict {
	case IceCubeTypeDry:
		return 10
	case IceCubeTypeWater:
		return 5
	default:
		panic("Invalid ice cube type")
	}
}

type iceCubes struct {
	cubesType IceCubeType
	quantity  uint
}

func NewIceCubes(cubesType IceCubeType, quantity uint) Condiment {
	return &iceCubes{
		cubesType: cubesType,
		quantity:  quantity,
	}
}

func (ic *iceCubes) String() string {
	return fmt.Sprintf(quantifiedDescriptionFormat, "Ice cubes", ic.quantity)
}

func (ic *iceCubes) GetCondimentCost() float64 {
	return ic.cubesType.GetCubeCost() * float64(ic.quantity)
}
