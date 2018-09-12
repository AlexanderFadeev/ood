package condiment

type lemon struct {
	condiment
}

func NewLemon(quantity uint) Condiment {
	return &lemon{
		*newQuantifiedCondiment("Lemon", 10, quantity),
	}
}
