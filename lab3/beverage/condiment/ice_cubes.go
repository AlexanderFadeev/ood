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
	condiment
}

func NewIceCubes(cubesType IceCubeType, quantity uint) Condiment {
	return &iceCubes{*newQuantifiedCondiment(
		fmt.Sprintf("%s ice cubes", cubesType),
		cubesType.GetCubeCost(), quantity,
	)}
}
