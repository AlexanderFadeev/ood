package condiment

import "fmt"

type chocolateCrumbs struct {
	weight float64
}

func NewChocolateCrumbs(weight uint) Condiment {
	return &chocolateCrumbs{
		weight: float64(weight),
	}
}

func (cc *chocolateCrumbs) String() string {
	return fmt.Sprintf(weightedDescriptionFormat, "Chocolate crumbs", cc.weight)
}

func (cc *chocolateCrumbs) GetCondimentCost() float64 {
	return 1 * cc.weight
}
