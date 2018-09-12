package condiment

type coconutFlakes struct {
	condiment
}

func NewCoconutFlakes(mass uint) Condiment {
	return &coconutFlakes{*newWeightedCondiment("Coconut flakes", 2, float64(mass))}
}
