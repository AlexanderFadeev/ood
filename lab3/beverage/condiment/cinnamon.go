package condiment

type cinnamon struct {
	condiment
}

func NewCinnamon() Condiment {
	return &cinnamon{*newCondiment("Cinnamon", 20)}
}
