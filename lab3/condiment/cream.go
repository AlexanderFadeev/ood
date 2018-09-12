package condiment

type cream struct {
	condiment
}

func NewCream() Condiment {
	return &cream{*newCondiment("Cream", 25)}
}
