package condiment

type cinnamon struct{}

func NewCinnamon() Condiment {
	return new(cinnamon)
}

func (cinnamon) String() string {
	return "Cinnamon"
}

func (cinnamon) GetCondimentCost() float64 {
	return 20
}
