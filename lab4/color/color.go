package color

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Color int

const (
	Red = Color(iota)
	Green
	Blue
	Pink
	Yellow
	Black
)

func (c *Color) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return errors.Wrap(err, "Failed to unmarshal color string")
	}

	switch str {
	case "red":
		*c = Red
	case "green":
		*c = Green
	case "blue":
		*c = Blue
	case "pink":
		*c = Pink
	case "yellow":
		*c = Yellow
	case "black":
		*c = Black
	default:
		return errors.Errorf("Invalid color `%s`", str)
	}

	return nil
}

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

func (c Color) String() string {
	switch c {
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Blue:
		return "Blue"
	case Pink:
		return "Pink"
	case Yellow:
		return "Yellow"
	case Black:
		return "Black"
	default:
		panic("Invalid color value")
	}
}
