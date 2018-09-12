package condiment

type chocolateCrumbs struct {
	condiment
}

func NewChocolateCrumbs(mass uint) Condiment {
	return &chocolateCrumbs{*newWeightedCondiment("Chocolate crumbs", 1, float64(mass))}
}
