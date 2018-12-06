package style

type lineStyle struct {
	styleImpl
	width float64
}

func NewLineStyle() LineStyle {
	return &lineStyle{
		styleImpl: newStyleImpl(),
		width:     1,
	}
}

func (ls *lineStyle) GetWidth() *float64 {
	return &ls.width
}

func (ls *lineStyle) SetWidth(width float64) {
	ls.width = width
}
