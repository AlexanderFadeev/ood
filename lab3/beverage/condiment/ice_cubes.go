package condiment

import "fmt"

type IceCubeType int

const (
	DryIce = IceCubeType(iota)
	WaterIce
)

func (ict IceCubeType) String() string {
	switch ict {
	case DryIce:
		return "Dry"
	case WaterIce:
		return "Water"
	default:
		panic("Invalid ice cube type")
	}
}

func (ict IceCubeType) GetCubeCost() float64 {
	switch ict {
	case DryIce:
		return 10
	case WaterIce:
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
