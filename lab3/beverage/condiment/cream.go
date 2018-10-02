package condiment

type cream struct{}

func NewCream() Condiment {
	return new(cream)
}

func (cream) String() string {
	return "Cream"
}

func (cream) GetCondimentCost() float64 {
	return 25
}
