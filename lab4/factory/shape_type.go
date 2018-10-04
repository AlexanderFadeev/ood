package factory

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type shapeType string

const (
	shapeTypeRectangle = shapeType(iota)
	shapeTypeTriangle
	shapeTypeEllipse
	shapeTypeRegularPolygon
)

func (st *shapeType) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return errors.Wrap(err, "Failed to unmarshal shape type string")
	}

	switch str {
	case "rectangle":
		*st = shapeTypeRectangle
	case "triangle":
		*st = shapeTypeTriangle
	case "ellipse":
		*st = shapeTypeEllipse
	case "regular polygon":
		*st = shapeTypeRegularPolygon
	default:
		return errors.Errorf("Invalid shape type `%s`", str)
	}

	return nil
}
