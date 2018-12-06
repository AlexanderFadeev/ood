package style

type fillStyle struct {
	styleImpl
}

func NewFillStyle() FillStyle {
	return &fillStyle{
		styleImpl: newStyleImpl(),
	}
}
