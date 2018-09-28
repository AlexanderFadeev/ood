package color

type Color int

const (
	Red = Color(iota)
	Green
	Blue
	Pink
	Yellow
	Black
)

func (c Color) RGBA() (r, g, b, a uint32) {
	switch c {
	case Red:
		return 0xFFFF, 0x0000, 0x0000, 0xFFFF
	case Green:
		return 0x0000, 0xFFFF, 0x0000, 0xFFFF
	case Blue:
		return 0x0000, 0x0000, 0xFFFF, 0xFFFF
	case Pink:
		return 0xFFFF, 0x7777, 0x7777, 0xFFFF
	case Yellow:
		return 0x7777, 0x7777, 0x0000, 0xFFFF
	case Black:
		return 0x0000, 0x0000, 0x0000, 0xFFFF
	default:
		panic("Invalid color value")
	}
}
