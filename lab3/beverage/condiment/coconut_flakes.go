package condiment

import "fmt"

type coconutFlakes struct {
	weight float64
}

func NewCoconutFlakes(weight float64) Condiment {
	return &coconutFlakes{
		weight: weight,
	}
}

func (cf *coconutFlakes) String() string {
	return fmt.Sprintf(weightedDescriptionFormat, "Coconut flakes", cf.weight)
}

func (cf *coconutFlakes) GetCondimentCost() float64 {
	return 2 * cf.weight
}
